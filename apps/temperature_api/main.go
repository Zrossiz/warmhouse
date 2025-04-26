package main

import (
	"encoding/json"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type responseDTO struct {
	Value int `json:"value"`
}

func main() {
	_ = godotenv.Load()

	if err := http.ListenAndServe(os.Getenv("ADDRESS"), http.HandlerFunc(handler)); err != nil {
		log.Fatalf("error starting web server: %v", err)
	}
}

func handler(rw http.ResponseWriter, r *http.Request) {
	path := r.URL
	isValidPattern := strings.Contains("/temperature", path.String())
	if isValidPattern && r.Method == "GET" {
		query := r.URL.Query()
		location := query.Get("location")
		sensorId := query.Get("sensorId")
		handleRoute(location, sensorId, rw)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusNotFound)
}

func handleRoute(location string, sensorID string, rw http.ResponseWriter) {
	data := responseDTO{
		Value: rand.IntN(100),
	}

	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}

	if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(rw).Encode(data); err != nil {
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
	}
	return
}
