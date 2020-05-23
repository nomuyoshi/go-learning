package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan string)
	for i := 0; i < 3; i++ {
		go fetch(queue)
	}

	queue <- "http://example1.com"
	queue <- "http://example2.com"
	queue <- "http://example3.com"
	queue <- "http://example4.com"
	queue <- "http://example5.com"
	queue <- "http://example6.com"

	close(queue)
	fmt.Println("All goroutines finished.")
}

func fetch(queue <-chan string) {
	for url := range queue {
		time.Sleep(10 * time.Second)
		fmt.Printf("%s finished!!\n", url)
	}
	return
}
