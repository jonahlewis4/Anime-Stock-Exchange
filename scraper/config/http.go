package config

import (
	"encoding/json"
	"net/http"
)

func (c *Config) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		c.HandleGet(w, r)
	} else if r.Method == "POST" {
		c.HandleSet(w, r)
	}
}

func (c *Config) HandleGet(w http.ResponseWriter, r *http.Request) {
	//encode the structure to json

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(c)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func (c *Config) HandleSet(w http.ResponseWriter, r *http.Request) {
	//decode the passed json
	decoder := json.NewDecoder(r.Body)
	//TODO make this thread safe
	err := decoder.Decode(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
