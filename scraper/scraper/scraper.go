package scraper

import (
	"fmt"
	"github.com/jonahlewis4/Anime-Stock-Exchange/scraper/config"
	"github.com/jonahlewis4/Anime-Stock-Exchange/scraper/server"
)

type Scraper struct {
	config config.Config
	server server.Restserver
}

func (s *Scraper) Init() {
	fmt.Println("Scraper starting")

	s.server.Init(&s.config)
}
