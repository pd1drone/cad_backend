package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func GetAgeRecommendation(db sqlx.Ext, age int, classification string) ([]string, error) {
	userAgeRecommendation := make([]string, 0)

	var recommendation string

	rows, err := db.Queryx(`SELECT Recommendations 
	FROM age_classification_recommendations 
	WHERE MinAge <= ? AND MaxAge >= ? AND 
	Classification = ?`,
		age, age, classification)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&recommendation)
		if err != nil {
			return nil, err
		}
		userAgeRecommendation = append(userAgeRecommendation, recommendation)
	}

	for _, a := range userAgeRecommendation {
		fmt.Println(a)
	}

	return userAgeRecommendation, nil
}

func GetBMIRecommendation(db sqlx.Ext, bmi float64, classification string) ([]string, error) {
	userBMIRecommendation := make([]string, 0)

	var recommendation string

	rows, err := db.Queryx(`SELECT Recommendations 
	FROM bmi_classification_recommendations 
	WHERE MinBMI <= ? AND MaxBMI >= ? AND 
	Classification = ?`,
		bmi, bmi, classification)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&recommendation)
		if err != nil {
			return nil, err
		}
		userBMIRecommendation = append(userBMIRecommendation, recommendation)
	}

	for _, a := range userBMIRecommendation {
		fmt.Println(a)
	}

	return userBMIRecommendation, nil
}

func GetHighBloodRecommendation(db sqlx.Ext, bloodpressure int, scenario string) ([]string, error) {
	userHighbloodRecommendation := make([]string, 0)

	var recommendation string

	rows, err := db.Queryx(`SELECT Recommendations 
	FROM highblood_pressure_recommendations 
	WHERE MinBloodPressure <= ? AND MaxBloodPressure >= ? AND 
	Scenario = ?`,
		bloodpressure, bloodpressure, scenario)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&recommendation)
		if err != nil {
			return nil, err
		}
		userHighbloodRecommendation = append(userHighbloodRecommendation, recommendation)
	}

	return userHighbloodRecommendation, nil
}

func SaveActiveRecommendations(db sqlx.Ext, accountID int, recommendation string) error {

	_, err := db.Exec(`INSERT INTO active_recommendations (
		AccountID,
		Recommendations
	)
	Values(?, ?)`,
		accountID,
		recommendation,
	)
	if err != nil {
		return err
	}

	return nil
}

func SaveActiveHighBloodRecommendations(db sqlx.Ext, accountID int, recommendation string) error {

	_, err := db.Exec(`INSERT INTO active_highblood_recommendations (
		AccountID,
		Recommendations
	)
	Values(?, ?)`,
		accountID,
		recommendation,
	)
	if err != nil {
		return err
	}

	return nil
}

func GetActiveRecommendations(db sqlx.Ext, accountid int) ([]string, error) {
	userActiveRecommendation := make([]string, 0)

	var recommendation string

	rows, err := db.Queryx(`SELECT Recommendations 
	FROM active_recommendations 
	WHERE AccountID = ? `,
		accountid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&recommendation)
		if err != nil {
			return nil, err
		}
		userActiveRecommendation = append(userActiveRecommendation, recommendation)
	}

	return userActiveRecommendation, nil
}

func GetActiveHighbloodRecommendations(db sqlx.Ext, accountid int) ([]string, error) {
	userActiveHighbloodRecommendation := make([]string, 0)

	var recommendation string

	rows, err := db.Queryx(`SELECT Recommendations 
	FROM active_highblood_recommendations 
	WHERE AccountID = ? `,
		accountid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&recommendation)
		if err != nil {
			return nil, err
		}
		userActiveHighbloodRecommendation = append(userActiveHighbloodRecommendation, recommendation)
	}

	return userActiveHighbloodRecommendation, nil
}

func DeleteActiveRecommendations(db sqlx.Ext, accountid int) error {

	_, err := db.Exec(`DELETE FROM active_recommendations WHERE AccountID = ? `, accountid)

	if err != nil {
		return err
	}

	return nil
}

func DeleteActiveHighbloodRecommendations(db sqlx.Ext, accountid int) error {

	_, err := db.Exec(`DELETE FROM active_highblood_recommendations WHERE AccountID = ? `, accountid)

	if err != nil {
		return err
	}

	return nil
}
