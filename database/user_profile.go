package database

import (
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type UserProfile struct {
	AccountID  int64  `json:"AccountID,omitempty"`
	Name       string `json:"Name,omitempty"`
	ProfileImg string `json:"ProfileImg,omitempty"`
}

type GetUserProfileResponse struct {
	UserProfiles []*UserProfile `json:"UserProfiles,omitempty"`
}

func GetUserProfile(db sqlx.Ext) (*GetUserProfileResponse, error) {
	userProfileArray := make([]*UserProfile, 0)

	var accountID int64
	var name string
	var img string

	rows, err := db.Queryx(`SELECT AccountID, Name, ProfileImg FROM user_profile`)
	if err != nil {
		return nil, fmt.Errorf("Error in retrieving user profile: %s", err)
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&accountID, &name, &img)
		if err != nil {
			return nil, err
		}
		userProfileArray = append(userProfileArray, &UserProfile{
			AccountID:  accountID,
			Name:       name,
			ProfileImg: img,
		})
	}

	return &GetUserProfileResponse{UserProfiles: userProfileArray}, nil
}

func AddUserProfile(db sqlx.Ext, profileName, profileImg string) (int64, error) {
	d, err := db.Exec(`INSERT INTO user_profile(
		Name,
		ProfileImg
	)
	Values(?,?)`,
		profileName,
		profileImg,
	)
	if err != nil {
		return 0, err
	}

	userID, err := d.LastInsertId()

	if err != nil {
		return 0, err
	}

	return userID, nil
}

func DeleteUserProfile(db sqlx.Ext, accountID int) error {
	_, err := db.Exec(`DELETE FROM user_profile WHERE AccountID = ?`, accountID)
	if err != nil {
		return err
	}

	_, err = db.Exec(`DELETE FROM user_details WHERE AccountID = ?`, accountID)
	if err != nil {
		return err
	}

	_, err = db.Exec(`DELETE FROM user_logs WHERE AccountID = ?`, accountID)
	if err != nil {
		return err
	}

	return nil
}

type UserDetailsResponse struct {
	AccountID  int    `json:"AccountID"`
	Name       string `json:"Name"`
	IsDiabetic bool   `json:"IsDiabetic"`
	Age        int    `json:"Age"`
	Weight     string `json:"Weight"`
	Height     string `json:"Height"`
	Gender     string `json:"Gender"`
}

func GetUserDetails(db sqlx.Ext, accountid int) (*UserDetailsResponse, error) {

	var AccountID int
	var name string
	var diabetic bool
	var age int
	var weight float64
	var height int
	var gender string

	rows, err := db.Queryx(`SELECT p.AccountID, p.Name,
    d.IsDiabetic, d.Age, d.Weight, d.Height, d.Gender
    FROM user_profile as p
    JOIN user_details as d
    ON  p.AccountID = d.AccountID
    WHERE p.AccountID = ?`, accountid)

	if err != nil {
		return nil, fmt.Errorf("Error in retrieving user details: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&AccountID, &name, &diabetic, &age, &weight, &height, &gender)
		if err != nil {
			return nil, err
		}

	}
	rows.Close()

	weightstring := fmt.Sprintf("%.2f", weight)
	heightstring := strconv.Itoa(height)

	return &UserDetailsResponse{
			AccountID:  accountid,
			Name:       name,
			IsDiabetic: diabetic,
			Age:        age,
			Weight:     weightstring,
			Height:     heightstring,
			Gender:     gender,
		},
		nil
}

type PostUserDetailsRequest struct {
	AccountID  int     `json:"AccountID"`
	Name       string  `json:"Name"`
	IsDiabetic bool    `json:"IsDiabetic"`
	Age        int     `json:"Age"`
	Weight     float64 `json:"Weight"`
	Height     int     `json:"Height"`
	Gender     string  `json:"Gender"`
}

func PostUserDetails(db sqlx.Ext, name string, diabetic bool, age int, weight float64, height int, gender string, accountid int) error {

	_, err := db.Exec(`UPDATE user_profile SET Name= ? WHERE AccountID= ?`, name, accountid)
	if err != nil {
		return err
	}

	var Count int

	Exists := false
	rows, err := db.Queryx(`SELECT COUNT(AccountID) FROM user_details WHERE AccountID = ?`, accountid)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&Count)
		if err != nil {
			return err
		}

		if Count > 0 {
			Exists = true
		} else {
			Exists = false
		}
	}
	rows.Close()

	if Exists {
		_, err := db.Exec(`UPDATE user_details SET IsDiabetic= ?,Age = ?, Weight=? , Height= ?,Gender = ?  WHERE AccountID= ?`,
			diabetic,
			age,
			weight,
			height,
			gender,
			accountid)
		if err != nil {
			return err
		}

		return nil
	}

	_, err = db.Exec(`INSERT INTO user_details(
		AccountID,
		IsDiabetic,
		Age,
		Weight,
		Height,
		Gender
	)
	Values(?,?,?,?,?,?)`,
		accountid,
		diabetic,
		age,
		weight,
		height,
		gender,
	)
	if err != nil {
		return err
	}

	return nil
}
