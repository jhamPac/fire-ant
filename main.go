package main

import (
	"fmt"
	"io/ioutil"
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
	fmt.Println("This is how you Makefile")
}
