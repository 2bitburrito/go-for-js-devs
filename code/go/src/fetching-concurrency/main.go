package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

func fetch(url string, wg *sync.WaitGroup, counter ...*atomic.Int32) bool {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	fmt.Println("Fetched: ", url, resp.StatusCode)
	if len(counter) > 0 && resp.StatusCode == 200 {
		counter[0].Add(1)
	}
	return resp.StatusCode == 200
}

func fetchURLsSequential(urls []string) int {
	count := 0
	for _, url := range urls {
		ok := fetch(url)
		if ok {
			count += 1
		}
	}
	return count
}

func fetchURLsAtomic(urls []string) int32 {
	var count atomic.Int32
	wg := sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &count)
	}
	wg.Wait()
	return count.Load()
}

func main() {
	urls := []string{
		"https://go.dev",
		"https://golang.org",
		"https://pkg.go.dev",
		"https://example.com",
		"https://httpbin.org/",
		"https://github.com",
	}

	start := time.Now()
	count := fetchURLsAtomic(urls)
	fmt.Println("Successful Fetches: ", count, "Out of: ", len(urls))
	fmt.Println("Time since start: ", time.Since(start))
}
