package main

import "fmt"

func main() {
	// dataはどこからでもアクセスできるが、loopData関数内のみからアクセスするように書かれている
	data := []int{10, 20, 30}
	// 送信専用チャネルを引数として受け取り
	// dataの値を1つずつチャネルに送信する
	loopData := func(ch chan<- int) {
		defer close(ch)
		for i := range data {
			ch <- data[i]
		}
	}

	ch := make(chan int)
	go loopData(ch)

	// チャネルから受信した値を出力
	for v := range ch {
		fmt.Println(v)
	}
	// 実行結果
	// 10
	// 20
	// 30
}
