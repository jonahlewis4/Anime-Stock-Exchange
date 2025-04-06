package app

import "net/http"

func (s *Scraper) HandleRun(w http.ResponseWriter, r *http.Request) {
	s.Scrape()
	w.WriteHeader(http.StatusOK)
}
