package web

import (
	"cad/database"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type UserHistoryRequest struct {
	AccountID int `json:"AccountID"`
}

type InsertUserHistoryRequest struct {
	AccountID      int    `json:"AccountID,omitempty"`
	GlucoseLevel   int    `json:"GlucoseLevel,omitempty"`
	Classification string `json:"Classification,omitempty"`
}

func (c *CADConfiguration) GetUserHistory(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	req := &UserHistoryRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	UserHistoryArr, err := database.GetUserHistory(c.cadDB, req.AccountID)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, UserHistoryArr)
}

func (c *CADConfiguration) AddUserHistory(w http.ResponseWriter, r *http.Request) {

	// Open the log file for appending
	logFile, err := os.OpenFile("/root/cad_backend/http_logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		respondJSON(w, 500, nil)
		return
	}
	defer logFile.Close()

	// Set the log output to the log file
	log.SetOutput(logFile)

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	req := &InsertUserHistoryRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Printf("Error: %s", err)
		respondJSON(w, 400, nil)
		return
	}

	err = database.AddUserHistory(c.cadDB, req.AccountID, req.GlucoseLevel, req.Classification)
	if err != nil {
		log.Printf("Error: %s", err)
		respondJSON(w, 400, nil)
		return
	}

	err = database.DeleteActiveRecommendations(c.cadDB, req.AccountID)
	if err != nil {
		log.Printf("Error: %s", err)
		respondJSON(w, 400, nil)
		return
	}

	err = database.DeleteActiveHighbloodRecommendations(c.cadDB, req.AccountID)
	if err != nil {
		log.Printf("Error: %s", err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, nil)
	return
}

func (c *CADConfiguration) GetLastUserHistory(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	req := &UserHistoryRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	LastUserHistory, err := database.GetLastUserHistory(c.cadDB, req.AccountID)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, LastUserHistory)
}
