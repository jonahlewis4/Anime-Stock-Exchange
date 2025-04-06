package app

import (
	"fmt"
	"net/http"
)

type Restserver struct{}

func (server *Restserver) Init(scraper *Scraper) {
	mux := http.NewServeMux()

	mux.HandleFunc("/config", scraper.Config.Handle)
	mux.HandleFunc("/scrape", scraper.HandleRun)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
	}
}
