package main

import (
	"log"
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

func main() {
	urls := []string{
		"https://www.youtube.com/",
		"a",
		"https://www.yahoo.co.jp/",
		"https://invalidhost",
	}

	for r := range checkStatus(urls) {
		if r.Error != nil {
			log.Printf("Error: %s\n", r.Error)
			continue
		}
		log.Println("Status: ", r.Response.Status)
	}
}

func checkStatus(urls []string) <-chan Result {
	result := make(chan Result)
	go func() {
		defer close(result)
		for _, url := range urls {
			res, err := http.Get(url)
			result <- Result{
				Error:    err,
				Response: res,
			}
		}
	}()

	return result
}
