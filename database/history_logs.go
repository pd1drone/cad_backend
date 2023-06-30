package database

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type AddGetUserHistoryData struct {
	AccountID      int    `json:"AccountID"`
	GlucoseLevel   int    `json:"GlucoseLevel"`
	Timestamp      int64  `json:"Timestamp"`
	Classification string `json:"Classification"`
}

type GetUserHistoryResponse struct {
	UserHistory []*AddGetUserHistoryData `json:"UserHistory"`
}

func AddUserHistory(db sqlx.Ext, accountID int, GlocuseLevel int, classification string) error {
	if accountID == 0 {
		return fmt.Errorf("Invalid Request")
	}
	_, err := db.Exec(`INSERT INTO user_logs(
		AccountID,
		GlucoseLevel,
		Timestamp,
		Classification
	)
	Values(?,?,?,?)`,
		accountID,
		GlocuseLevel,
		time.Now().Unix(),
		classification,
	)
	if err != nil {
		return err
	}

	return nil
}

func GetUserHistory(db sqlx.Ext, accountid int) (*GetUserHistoryResponse, error) {
	userHistoryArray := make([]*AddGetUserHistoryData, 0)

	var accountID int64
	var glocose int
	var timestamp int64
	var classification string

	rows, err := db.Queryx(`SELECT AccountID, GlucoseLevel,Timestamp, Classification FROM user_logs WHERE AccountID = ? ORDER BY Timestamp DESC`, accountid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&accountID, &glocose, &timestamp, &classification)
		if err != nil {
			return nil, err
		}
		userHistoryArray = append(userHistoryArray, &AddGetUserHistoryData{
			AccountID:      accountid,
			GlucoseLevel:   glocose,
			Timestamp:      timestamp,
			Classification: classification,
		})
	}

	return &GetUserHistoryResponse{UserHistory: userHistoryArray}, nil
}

func GetLastUserHistory(db sqlx.Ext, accountid int) (*AddGetUserHistoryData, error) {

	var accountID int64
	var glocose int
	var timestamp int64
	var classification string

	rows, err := db.Queryx(`SELECT AccountID, GlucoseLevel,Timestamp, Classification FROM user_logs WHERE AccountID = ? ORDER BY Timestamp DESC LIMIT 1`, accountid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&accountID, &glocose, &timestamp, &classification)
		if err != nil {
			return nil, err
		}
	}

	return &AddGetUserHistoryData{
		AccountID:      accountid,
		GlucoseLevel:   glocose,
		Timestamp:      timestamp,
		Classification: classification,
	}, nil
}
