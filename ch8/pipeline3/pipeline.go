package main

import "fmt"

// チャネルを使って3つのgoroutineを接続する
// 1つ目(Counter)、整数を0,1,2...と生成し、2つ目のgoroutineに送信
// 2つ目(Squarer)、二乗して、その結果を3つ目のgoroutineに送信
// 3つ目(Printer)、受信した値を表示する

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)

	printer(squares)
}
