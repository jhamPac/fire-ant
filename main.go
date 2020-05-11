package main

import (
	"fmt"
	"net/http"

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

	doc.Find(".div-col").Each
}

func main() {
	url := "https://en.wikipedia.org/wiki/The_Pitchfork_500"

	body := parseWiki(url)

	fmt.Println(body)
}
