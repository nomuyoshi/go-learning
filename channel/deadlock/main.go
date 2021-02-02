package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		// デッドロックを起こす
		// 必ずreturn。チャネルに送信することなく、ゴルーチンは終了する
		if true {
			fmt.Println("goroutine1 done")
			return
		}
		// チャネルに値を送信
		ch <- "Hello, Channel!!"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("goroutine2 done")
	}()

	// チャネルから値を受信
	// 2つのゴルーチンが終了したとき、チャネルに送信するゴルーチンが存在しないためデッドロックが起こる
	fmt.Println(<-ch)
}
