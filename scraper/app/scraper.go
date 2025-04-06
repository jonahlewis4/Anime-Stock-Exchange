package app

import (
	"encoding/json"
	"fmt"
	config "github.com/jonahlewis4/Anime-Stock-Exchange/scraper/app/config"
	api "github.com/jonahlewis4/Anime-Stock-Exchange/scraper/app/response"
	"net/http"
	"strings"
)

type Scraper struct {
	Config config.Config
	server Restserver
}

func (s *Scraper) Init() {
	fmt.Println("Scraper starting")

	s.server.Init(s)
}

func (s *Scraper) Scrape() {
	fmt.Println("Scraping...")
	//TODO get the info

	request, err := http.NewRequest("GET", "https://www.animecharactersdatabase.com/api_series_characters.php?character_id=56145", strings.NewReader(""))
	if err != nil {
		fmt.Println("error creating request", err)
		return
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("Error sending request", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error hitting 3rd party character api: ", response.Status)
	}

	apiResponse := &api.Response{}
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(apiResponse)
	if err != nil {
		fmt.Println("Error decoding api response", err)
	}

	fmt.Println(apiResponse)

}
