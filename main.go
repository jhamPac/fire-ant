package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func makeRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	return string(body), nil
}

func main() {
	url := "https://en.wikipedia.org/wiki/The_Pitchfork_500"

	body, err := makeRequest(url)
	if err != nil {
		log.Fatalf("Request failed %v", err)
	}

	fmt.Println(body)
}
