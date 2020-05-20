package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	queue := make(chan string)
	// 3つのゴルーチンを生成
	for i := 0; i < 3; i++ {
		wg.Add(1) // ゴルーチン生成前にAddでインクリメントする
		go fetch(queue)
	}

	queue <- "http://example1.com"
	queue <- "http://example2.com"
	queue <- "http://example3.com"
	queue <- "http://example4.com"
	queue <- "http://example5.com"
	queue <- "http://example6.com"

	close(queue)
	wg.Wait()
	fmt.Println("All goroutines finished.")
}

func fetch(queue <-chan string) {
	for url := range queue {
		time.Sleep(10 * time.Second)
		fmt.Printf("%s finished!!\n", url)
	}
	// 処理が完了したらDoneでデクリメントする
	wg.Done()
	return
}
