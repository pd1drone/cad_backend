package web

import (
	"cad/database"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ClassifyRequest struct {
	AccountID    int `json:"AccountID"`
	GlucoseLevel int `json:"GlucoseLevel"`
}

type ClassifyResponse struct {
	Classification string `json:"Classification"`
}

func (c *CADConfiguration) ClassifyGlocuseLevel(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	req := &ClassifyRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	userDetails, err := database.GetUserDetails(c.cadDB, req.AccountID)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	classification, err := database.ClassifyGlocuseLevel(c.cadDB, userDetails.Age, req.GlucoseLevel)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, ClassifyResponse{
		Classification: classification,
	})
}
