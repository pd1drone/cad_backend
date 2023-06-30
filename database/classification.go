package database

import (
	"github.com/jmoiron/sqlx"
)

func ClassifyGlocuseLevel(db sqlx.Ext, age int, glocuseLevel int) (string, error) {

	var classification string

	rows, err := db.Queryx(`SELECT Classification 
	FROM classify 
	WHERE MinAge <= ? AND MaxAge >= ? AND 
	MinGlocuseLevel <= ? AND MaxGlocuseLevel >= ?`,
		age, age, glocuseLevel, glocuseLevel)
	if err != nil {
		return "", err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&classification)
		if err != nil {
			return "", err
		}
	}

	return classification, nil
}
