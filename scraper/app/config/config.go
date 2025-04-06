package scraper

import (
	"time"
)

type Config struct {
	On                   bool          `json:"On"`                   //on indicates if the scraper is on or not
	MillisBetweenScrapes time.Duration `json:"MillisBetweenScrapes"` //if the scraper is on, it will scrape every interval
}

func (c *Config) Init() {
	//TODO get from a dabatase
	*c = Config{
		On:                   false,
		MillisBetweenScrapes: time.Hour * 24,
	}
}
