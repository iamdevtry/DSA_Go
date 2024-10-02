package main

import (
	"fmt"
	"time"
)

const (
	maxNumbUrls = 10000
	maxWorkers  = 5
)

func queueUrls() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i <= maxNumbUrls; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func crawl(ch <-chan int, name string) {
	for v := range ch {
		fmt.Printf("%s crawling %v\n", name, v)
		time.Sleep(time.Second)
	}
}

func main() {
	urls := queueUrls()

	for i := 0; i < maxWorkers; i++ {
		go crawl(urls, fmt.Sprintf("worker %v", i))
	}

	time.Sleep(1 * time.Minute)
}
