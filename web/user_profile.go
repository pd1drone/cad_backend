package web

import (
	"cad/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AddDeleteUserProfileRequest struct {
	AccountID  int    `json:"AccountID,omitempty"`
	Name       string `json:"Name,omitempty"`
	ProfileImg string `json:"ProfileImg,omitempty"`
}

type AddUserProfileResponse struct {
	AccountID int `json:"AccountID"`
}

func (c *CADConfiguration) GetUserProfile(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	UserProfileArr, err := database.GetUserProfile(c.cadDB)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, UserProfileArr)
}

func (c *CADConfiguration) AddUserProfile(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	req := &AddDeleteUserProfileRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	accntID, err := database.AddUserProfile(c.cadDB, req.Name, req.ProfileImg)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, AddUserProfileResponse{
		AccountID: int(accntID),
	})
}

func (c *CADConfiguration) DeleteUserProfile(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	req := &AddDeleteUserProfileRequest{}

	fmt.Println(req.AccountID)

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	err = database.DeleteUserProfile(c.cadDB, req.AccountID)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, nil)
}

func (c *CADConfiguration) GetUserDetails(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	req := &AddDeleteUserProfileRequest{}

	fmt.Println(req.AccountID)

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	UserDetails, err := database.GetUserDetails(c.cadDB, req.AccountID)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, UserDetails)
}

func (c *CADConfiguration) PostUserDetails(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	req := &database.PostUserDetailsRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	err = database.PostUserDetails(c.cadDB, req.Name, req.IsDiabetic, req.Age, req.Weight, req.Height, req.Gender, req.AccountID)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, nil)
}
