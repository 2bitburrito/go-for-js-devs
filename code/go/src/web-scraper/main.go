package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func fetchURL(url string, ch chan bool) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("✗ %s: %v\n", url, err)
		ch <- false
		return
	}
	defer resp.Body.Close()

	ch <- resp.StatusCode == 200
}

func main() {
	var urls []string
	jsonFileBytes, err := os.ReadFile("../../../shared/urls.json")
	if err != nil {
		panic("couldn't read file: " + err.Error())
	}

	err = json.Unmarshal(jsonFileBytes, &urls)
	if err != nil {
		panic("couldn't decode json: " + err.Error())
	}

	start := time.Now()

	ch := make(chan bool)

	for _, url := range urls {
		go fetchURL(url, ch)
	}

	count := 0
	for range len(urls) {
		if <-ch {
			count++
		}
	}
	close(ch)

	fmt.Printf("Scraped a total of %d URLs of %d\n", count, len(urls))
	fmt.Printf("Total time: %v\n", time.Since(start))
}
