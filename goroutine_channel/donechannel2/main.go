package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 値を送信する処理をdoneチャネルを使ってキャンセルする例
func main() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream 終了")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					fmt.Println("newRandStream キャンセル")
					return
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)
	//newRandStreamのゴルーチンの出力を確認するためにSleepしてメインゴルーチンが終了しないようにする
	time.Sleep(1 * time.Second)
}
