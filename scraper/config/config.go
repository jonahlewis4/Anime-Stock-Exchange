package config

import (
	"encoding/json"
	"net/http"
	"time"
)

type Config struct {
	On                   bool          `json:"On"`                   //on indicates if the scraper is on or not
	MillisBetweenScrapes time.Duration `json:"MillisBetweenScrapes"` //if the scraper is on, it will scrape every interval
}

var _config = Config{
	On:                   false,
	MillisBetweenScrapes: time.Hour * 24,
}

func HandleConfig(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		HandleGetConfig(w, r)
	} else if r.Method == "POST" {
		HandleSetConfig(w, r)
	}
}

func HandleGetConfig(w http.ResponseWriter, r *http.Request) {
	//encode the structure to json

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&_config)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func HandleSetConfig(w http.ResponseWriter, r *http.Request) {
	//decode the passed json
	decoder := json.NewDecoder(r.Body)
	//TODO make this thread safe
	err := decoder.Decode(&_config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
