package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "A"
		close(ch2)
	}()

	value1, ok1 := <-ch1
	value2, ok2 := <-ch2
	fmt.Printf("value1: %s, ok1: %t\n", value1, ok1) // value1: A, ok1: true
	fmt.Printf("value2: %s, ok2: %t\n", value2, ok2) // value2: , ok2: false

	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
