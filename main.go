package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func parseWiki(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET request failed: %v", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("goquery could not read the body: %v", err)
	}

	var tracks []string
	var track string

	tracks = append(tracks, "artist,song\n")

	doc.Find(".div-col").Each(func(_ int, s *goquery.Selection) {
		s.Find("li").Each(func(_ int, t *goquery.Selection) {
			text := strings.Split(t.Text(), " –")

			artist := text[0]
			song := strings.Trim(text[1], " \"")

			track = artist + "," + song + "\n"

			tracks = append(tracks, track)
		})
	})

	return tracks, nil
}

func main() {
	url := "https://en.wikipedia.org/wiki/The_Pitchfork_500"

	body, err := parseWiki(url)
	if err != nil {
		log.Fatal("Parsing the url failed sorry")
	}

	fmt.Println(body)
}
