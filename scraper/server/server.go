package server

import (
	"fmt"
	"github.com/jonahlewis4/Anime-Stock-Exchange/scraper/config"
	"net/http"
)

type Restserver struct{}

func (server *Restserver) Init() {
	mux := http.NewServeMux()

	mux.HandleFunc("/config", config.HandleConfig)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
	}
}
