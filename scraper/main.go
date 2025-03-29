package main

import (
	"fmt"
	"github.com/jonahlewis4/Anime-Stock-Exchange/scraper/config"
	"net/http"
)

func main() {
	fmt.Println("Scraper started")

	mux := http.NewServeMux()
	
	mux.HandleFunc("/config", config.HandleConfig)

}
