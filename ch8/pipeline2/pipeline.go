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
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		// チャネルを閉じる。閉じることで、受信側は待たされることなく、ゼロ値が生成される。
		// 閉じられてチャネルに送信しようするとpanicになる。
		close(naturals)
	}()

	// Squarer
	go func() {
		for {
			// チャネルnaturalsから整数を受信
			// ok: 成功したらtrue, 閉じられて空になったチャネルからの受信はfalse
			x, ok := <-naturals
			if !ok {
				break
			}
			// 受信した値を二乗して、チャネルsquaresに送信
			squares <- x * x
		}
		close(squares)
	}()

	// Printer メインゴルーチン
	for {
		// チャネルの受信にもrangeループが使える。
		// チャネルに対して送信されたすべての値を受信し、最後の値の後にループを終了する
		for x := range squares {
			fmt.Println(x)
		}
	}
}
