package config

import (
	"encoding/json"
	"net/http"
	"time"
)

type config struct {
	On                 bool          `json:"On"`                 //on indicates if the scraper is on or not
	TimeBetweenScrapes time.Duration `json:"TimeBetweenScrapes"` //if the scraper is on, it will scrape every interval
}

var _config = config{}

func HandleConfig(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		HandleGetConfig(w, r)
	} else if r.Method == "POST" {
		HandleSetConfig(w, r)
	}
}

func HandleGetConfig(w http.ResponseWriter, r *http.Request) {
	//encode the structure to json
	encoder := json.NewEncoder(w)
	err := encoder.Encode(w)
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
