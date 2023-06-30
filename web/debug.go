package web

import (
	"cad/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DebugAnalogVoltageRequest struct {
	VoltageData float64 `json:"VoltageData"`
}

func (c *CADConfiguration) DebugAnalogVoltage(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	req := &DebugAnalogVoltageRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println(err)
		respondJSON(w, 400, nil)
		return
	}

	err = database.DebugAnalogVoltageValue(c.cadDB, req.VoltageData)
	if err != nil {
		fmt.Println(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, nil)
}
