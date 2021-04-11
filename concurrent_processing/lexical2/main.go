package main

import (
	"fmt"
	"sync"
)

func main() {
	sumPrinter := func(wg *sync.WaitGroup, data []int) {
		defer wg.Done()

		var sum int
		for _, v := range data {
			sum += v
		}

		fmt.Printf("sum = %d\n", sum)
	}
	var wg sync.WaitGroup
	data := []int{1, 2, 3, 10, 20, 30}

	wg.Add(2)
	go sumPrinter(&wg, data[3:]) // => wum = 6
	go sumPrinter(&wg, data[:3]) // => sum = 60
	wg.Wait()

	fmt.Println("Done!!")
}
