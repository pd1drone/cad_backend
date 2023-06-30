package web

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (c *CADConfiguration) check(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func Routes() {
	// Check if the log file exists
	if _, err := os.Stat("/root/cad_backend/http_logs.log"); os.IsNotExist(err) {
		// Create a log file
		logFile, err := os.Create("/root/cad_backend/http_logs.log")
		if err != nil {
			log.Fatal(err)
		}
		defer logFile.Close()
	}

	// Open the log file
	logFile, err := os.OpenFile("/root/cad_backend/http_logs.log", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// // Create a new logger using the log file
	logger := middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger:  log.New(logFile, "", log.LstdFlags), // Use the log file as the output
		NoColor: true,
	})

	log.Print("Starting CAD Backend Service.....")
	r := chi.NewRouter()

	cadCfg, err := New()
	if err != nil {
		log.Fatal(err)
	}
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(logger)
	r.Use(func(next http.Handler) http.Handler {
		return logPayloads(logFile, next)
	})

	r.Get("/check", cadCfg.check)
	r.Get("/getUserProfile", cadCfg.GetUserProfile)
	r.Post("/addUserProfile", cadCfg.AddUserProfile)
	r.Post("/deleteUserProfile", cadCfg.DeleteUserProfile)
	r.Post("/getUserDetails", cadCfg.GetUserDetails)
	r.Post("/postUserDetails", cadCfg.PostUserDetails)
	r.Post("/getLastUserHistory", cadCfg.GetLastUserHistory)
	r.Post("/addUserHistory", cadCfg.AddUserHistory)
	r.Post("/getUserHistory", cadCfg.GetUserHistory)
	r.Post("/login", cadCfg.LoginUser)
	r.Get("/getActiveUser", cadCfg.GetActiveUser)
	r.Post("/getRecommendations", cadCfg.GetRecommendations)
	r.Post("/getActiveRecommendations", cadCfg.GetActiveRecommendations)
	r.Post("/sendDebugAnalogVoltage", cadCfg.DebugAnalogVoltage)
	r.Post("/classify", cadCfg.ClassifyGlocuseLevel)
	r.Get("/logout", cadCfg.LogOutUser)

	http.ListenAndServe("0.0.0.0:8081", r)
}

func logPayloads(logFile *os.File, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Record start time of request processing
		startTime := time.Now()

		// Call the next middleware/handler in the chain and capture response
		responseRecorder := httptest.NewRecorder()
		next.ServeHTTP(responseRecorder, r)

		// Record end time of request processing
		endTime := time.Now()

		// Read request payload
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading request body:", err)
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		// Restore request body after reading
		r.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

		// Read response payload
		responseBody := responseRecorder.Body.Bytes()

		// Log the request and response payloads
		log.SetOutput(logFile)
		log.Printf("Request: %s %s\nPayload: %s\nResponse: %d %s\nPayload: %s\nDuration: %v\n",
			r.Method, r.URL.Path, string(requestBody), responseRecorder.Code, http.StatusText(responseRecorder.Code),
			string(responseBody), endTime.Sub(startTime))

		// Write response back to original response writer
		for k, v := range responseRecorder.Header() {
			w.Header()[k] = v
		}
		w.WriteHeader(responseRecorder.Code)
		w.Write(responseBody)
	})
}
