package database

import "github.com/jmoiron/sqlx"

type ActiveUserResponse struct {
	AccountID int `json:"AccountID"`
}

func GetActiveUser(db sqlx.Ext) (*ActiveUserResponse, error) {

	var accountID int64

	rows, err := db.Queryx(`SELECT AccountID FROM active_user`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&accountID)
		if err != nil {
			return nil, err
		}
	}

	return &ActiveUserResponse{AccountID: int(accountID)}, nil
}

func LoginUser(db sqlx.Ext, accountID int) error {

	_, err := db.Exec(`INSERT INTO active_user (
		AccountID
	)
	Values(?)`,
		accountID,
	)
	if err != nil {
		return err
	}

	return nil
}

func LogOutUser(db sqlx.Ext) error {

	_, err := db.Exec(`TRUNCATE TABLE active_user`)

	if err != nil {
		return err
	}

	return nil
}
