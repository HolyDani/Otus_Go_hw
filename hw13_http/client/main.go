package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	var url, path string
	flag.StringVar(&url, "url", "https://www.google.com/", "URL")
	flag.StringVar(&path, "path", "/", "Endpoint")
	flag.Parse()

	fullURL := url + path

	err := getRequest(fullURL)
	if err != nil {
		fmt.Println(err)
	}
	err = postRequest(fullURL)
	if err != nil {
		fmt.Println(err)
	}
}

func getRequest(url string) error {
	resp, err := http.Get(url) //nolint
	if err != nil {
		return fmt.Errorf("unable request %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read %w", err)
	}
	fmt.Println(resp.Status)
	fmt.Println(string(body))
	return nil
}

func postRequest(url string) error {
	resp, err := http.Post(url, "text", strings.NewReader("Daniel Vladimirov")) //nolint
	if err != nil {
		return fmt.Errorf("unable request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read %w", err)
	}
	fmt.Println(resp.Status)
	fmt.Println(string(body))
	return nil
}
