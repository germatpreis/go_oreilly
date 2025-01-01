package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func fetchUrl(url string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprintf("Error fetching %s: %w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error reading respone from %s: %s", url, string(body))
	}

	ch <- fmt.Sprintf("Response from %s: %s", url, string(body))
}

func main() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))

	ch := make(chan string, len(urls))

	for _, url := range urls {
		go fetchUrl(url, &wg, ch)
	}

	wg.Wait()
	close(ch)

	for response := range ch {
		fmt.Println(response)
	}
}
