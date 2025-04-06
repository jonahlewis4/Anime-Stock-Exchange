package app

import (
	"fmt"
	config "github.com/jonahlewis4/Anime-Stock-Exchange/scraper/app/config"
)

type Scraper struct {
	Config config.Config
	server Restserver
}

type Response struct {
}

func (s *Scraper) Init() {
	fmt.Println("Scraper starting")

	s.server.Init(s)
}

func (s *Scraper) Scrape() {
	fmt.Println("Scraping...")
	//TODO get the info
}
