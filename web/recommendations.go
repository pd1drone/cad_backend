package web

import (
	"cad/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type GetActiveRecommendationRequest struct {
	AccountID int `json:"AccountID"`
}
type RecommendationRequest struct {
	AccountID      int     `json:"AccountID"`
	Age            int     `json:"Age"`
	BMI            float64 `json:"BMI"`
	Classification string  `json:"Classification"`
	IsHighBlood    bool    `json:"IsHighBlood"`
	BloodPressure  int     `json:"BloodPressure"`
	Scenario       string  `json:"Scenario"`
}

type RecommendationResponse struct {
	Recommendations          []string `json:"Recommendations"`
	HighBloodRecommendations []string `json:"HighBloodRecommendations"`
}

func GenerateTwoRandomIndex(lengthOfArray int) (firstIndex, SecondIndex int) {
	//initialize firstIndex random generator
	rand.Seed(time.Now().Unix())
	firstIndex = rand.Intn(lengthOfArray)
	//initialize secondIndex random generator
	rand.Seed(time.Now().Unix() + int64(firstIndex))
	secondIndex := rand.Intn(lengthOfArray)
	// check index random generator
	for secondIndex == firstIndex {
		rand.Seed(time.Now().Unix() + int64(secondIndex))
		secondIndex = rand.Intn(lengthOfArray)
		if secondIndex != firstIndex {
			break
		}
	}

	return firstIndex, secondIndex

}

func (c *CADConfiguration) GetRecommendations(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	req := &RecommendationRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	ageRecommendation, err := database.GetAgeRecommendation(c.cadDB, req.Age, req.Classification)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	recommendation := make([]string, 0)
	highbloodRecommendation := make([]string, 0)
	index1, index2 := GenerateTwoRandomIndex(len(ageRecommendation) - 1)

	fmt.Println(index1)
	fmt.Println(index2)
	if len(ageRecommendation) >= 2 {
		recommendation = append(recommendation, ageRecommendation[index1], ageRecommendation[index2])
	} else {
		recommendation = append(recommendation, ageRecommendation...)
	}

	for _, d := range recommendation {
		fmt.Println(d)
	}

	bmiRecommendation, err := database.GetBMIRecommendation(c.cadDB, req.BMI, req.Classification)
	if err != nil {
		fmt.Println(err)
		respondJSON(w, 400, nil)
		return
	}

	bindex1, bindex2 := GenerateTwoRandomIndex(len(bmiRecommendation) - 1)

	fmt.Println(bindex1)
	fmt.Println(bindex2)
	if len(bmiRecommendation) >= 2 {
		recommendation = append(recommendation, bmiRecommendation[bindex1], bmiRecommendation[bindex2])
	} else {
		recommendation = append(recommendation, bmiRecommendation...)
	}

	for _, d := range recommendation {
		fmt.Println(d)
	}

	for _, row := range recommendation {
		err := database.SaveActiveRecommendations(c.cadDB, req.AccountID, row)
		if err != nil {
			respondJSON(w, 400, nil)
			return
		}
	}

	if req.IsHighBlood {
		recommendationHighblood, err := database.GetHighBloodRecommendation(c.cadDB, req.BloodPressure, req.Scenario)
		if err != nil {
			respondJSON(w, 400, nil)
			return
		}
		highbloodRecommendation = recommendationHighblood

		for _, row := range highbloodRecommendation {
			err := database.SaveActiveHighBloodRecommendations(c.cadDB, req.AccountID, row)
			if err != nil {
				respondJSON(w, 400, nil)
				return
			}
		}
	}

	respondJSON(w, 200, RecommendationResponse{
		Recommendations:          recommendation,
		HighBloodRecommendations: highbloodRecommendation,
	})
}

func (c *CADConfiguration) GetActiveRecommendations(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	req := &GetActiveRecommendationRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	activeRecommendation, err := database.GetActiveRecommendations(c.cadDB, req.AccountID)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	activeHighbloodRecommendation, err := database.GetActiveHighbloodRecommendations(c.cadDB, req.AccountID)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, &RecommendationResponse{
		Recommendations:          activeRecommendation,
		HighBloodRecommendations: activeHighbloodRecommendation,
	})
}
