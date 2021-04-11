package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan struct{})
	go func() {
		time.Sleep(2 * time.Second)
		close(ch)
	}()

	fmt.Println("Selectブロック中")
	select {
	case <-ch:
		fmt.Printf("ch1 close. ブロック時間: %v\n", time.Since(start))
	}
}
