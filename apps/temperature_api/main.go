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
	if strings.Contains("/temperature", path.String()) {
		data := responseDTO{
			Value: rand.IntN(100),
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(rw).Encode(data); err != nil {
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusNotFound)
}
