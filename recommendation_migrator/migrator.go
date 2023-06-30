package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gopkg.in/ini.v1"
)

func InitializeKioskDatabase(dbname, username, password, dbhost, dbport string) (*sqlx.DB, error) {
	conn := username + ":" + password + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname
	caddb, err := sqlx.Connect("mysql", conn)

	if err != nil {
		return nil, fmt.Errorf("Error in initializing cad database: %s", err)
	}

	return caddb, nil
}

type CADConfiguration struct {
	cadDB *sqlx.DB
}

func New() (*CADConfiguration, error) {
	// read config file
	cfg, err := ini.Load("config.ini")
	if err != nil {
		return nil, fmt.Errorf("Fail to read file: %v", err)
	}

	dbSection := cfg.Section("db")
	user := dbSection.Key("user").String()
	password := dbSection.Key("password").String()
	dbhost := dbSection.Key("dbhost").String()
	dbport := dbSection.Key("dbport").String()
	dbname := dbSection.Key("dbname").String()

	caddb, err := InitializeKioskDatabase(dbname, user, password, dbhost, dbport)
	if err != nil {
		return nil, err
	}

	return &CADConfiguration{
		caddb,
	}, nil
}

func main() {
	// Parse command line arguments
	csvFilePtr := flag.String("csv-file", "", "Path to CSV file")
	flag.Parse()

	table := ""
	switch *csvFilePtr {
	case "age_recommendations.csv":
		table = "age_classification_recommendations"
	case "bmi_recommendations.csv":
		table = "bmi_classification_recommendations"
	case "highblood_pressure_recommendation.csv":
		table = "highblood_pressure_recommendations"
	default:
		table = ""
	}

	fmt.Println(table)

	if table == "" {
		panic(fmt.Errorf("wrong csv used"))
	}

	cadCFG, err := New()
	if err != nil {
		log.Fatal(err)
	}

	// Open the CSV file
	file, err := os.Open(*csvFilePtr)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the CSV file into a slice of slices
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Get the headers (the first row of the CSV)
	headers := rows[0]
	columns := strings.Join(headers, ",")
	fmt.Println(columns)

	// Get the data (all the rows after the first)
	data := rows[1:]

	for _, row := range data {
		values := strings.Join(row, ",")
		query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, columns, values)
		fmt.Println(query)
		// Call your function to save the data to the database
		err := saveToDB(*cadCFG.cadDB, query)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func saveToDB(db sqlx.DB, query string) error {
	_, err := db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}
