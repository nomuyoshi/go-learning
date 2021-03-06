package main

import "fmt"

// チャネルを使って3つのgoroutineを接続する
// 1つ目(Counter)、整数を0,1,2...と生成し、2つ目のgoroutineに送信
// 2つ目(Squarer)、二乗して、その結果を3つ目のgoroutineに送信
// 3つ目(Printer)、受信した値を表示する

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		// 無限にインクリメントした整数をチャネルに送信
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			// チャネルnaturalsから整数を受信
			x := <-naturals
			// 受信した値を二乗して、チャネルsquaresに送信
			squares <- x * x
		}
	}()

	// Printer メインゴルーチン
	for {
		// チャネルsquaresから整数を受信して、Printlnにわたす
		fmt.Println(<-squares)
	}
}
