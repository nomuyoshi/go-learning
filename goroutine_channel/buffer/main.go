package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Printf("sent %d\n", i)
		}
	}()

	time.Sleep(3 * time.Second)
	for v := range ch {
		fmt.Printf("received value: %d\n", v)
		time.Sleep(1 * time.Second)
	}
}
