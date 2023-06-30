package web

import (
	"cad/database"
	"fmt"

	"github.com/jmoiron/sqlx"
	"gopkg.in/ini.v1"
)

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

	caddb, err := database.InitializeKioskDatabase(dbname, user, password, dbhost, dbport)
	if err != nil {
		return nil, err
	}

	return &CADConfiguration{
		caddb,
	}, nil
}
