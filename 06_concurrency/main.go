package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func sites() []string {
	return []string{
		"https://google.com",
		"https://reddit.com",
		"https://twitter.com",
		"https://as.com",
		"https://sport.es",
		"https://cse.gob.ni",
		"https://uam.edu.ni",
		"https://nfl.com",
		"https://nba.com",
		"https://go.dev",
	}
}

func makeRequest(url string) {
	now := time.Now()
	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("could not download site: %s \n", url)
	}

	fmt.Printf("took %s to download %s \n", time.Since(now), url)
}

func getSequencial() {
	list := sites()
	now := time.Now()
	for i := 0; i < len(list); i++ {
		makeRequest(list[i])
	}
	fmt.Println("---")
	fmt.Printf("took %s to download %d sites", time.Since(now), len(list))
}

func getConcurrent() {
	list := sites()
	now := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < len(list); i++ {
		wg.Add(1)
		go func(index int) {
			makeRequest(list[index])
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("---")
	fmt.Printf("took %s to download %d sites", time.Since(now), len(list))
}

func main() {
	// getSequencial()
	getConcurrent()
}
