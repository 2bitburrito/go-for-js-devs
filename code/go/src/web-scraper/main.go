package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func fetchURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("✗ %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	io.ReadAll(resp.Body) // Read the body
	fmt.Printf("Done! %s (%v)\n", url, time.Since(start))
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

	fmt.Println("Go Concurrent Scraping:")
	start := time.Now()

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, &wg)
	}

	wg.Wait()
	fmt.Printf("Total time: %v\n", time.Since(start))
}
