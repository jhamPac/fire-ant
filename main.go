package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func parseWiki(url string) string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	var tracks []string
	var track string

	tracks = append(tracks, "artist,song\n")
}

func main() {
	url := "https://en.wikipedia.org/wiki/The_Pitchfork_500"

	body := parseWiki(url)

	fmt.Println(body)
}
