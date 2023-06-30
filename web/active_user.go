package web

import (
	"cad/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *CADConfiguration) GetActiveUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	UserAccount, err := database.GetActiveUser(c.cadDB)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}
	if UserAccount.AccountID == 0 {
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, UserAccount)
}

func (c *CADConfiguration) LoginUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	req := &database.ActiveUserResponse{}

	fmt.Println(req.AccountID)

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	err = database.LoginUser(c.cadDB, req.AccountID)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, nil)
}

func (c *CADConfiguration) LogOutUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	err := database.LogOutUser(c.cadDB)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, nil)
}
