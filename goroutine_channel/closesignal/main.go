package main

import (
	"fmt"
	"sync"
)

func main() {
	begin := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		// ループでゴルーチンを5つ実行
		go func(i int) {
			defer wg.Done()
			// チャネルを受信
			<-begin
			// なにか処理をする
			fmt.Printf("%d do somthing...\n", i)
		}(i)
	}

	fmt.Println("close channel")
	close(begin)
	wg.Wait()
	// ↓↓実行結果↓↓

	// close channel
	// 1 do somthing...
	// 0 do somthing...
	// 2 do somthing...
	// 3 do somthing...
	// 4 do somthing...
}
