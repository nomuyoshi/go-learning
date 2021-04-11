package main

import "fmt"

func main() {
	chanOwner := func() <-chan int {
		// 関数内でチャネルを初期化して、受信操作専用のチャネルを返す
		// こうすることでチャネルへ書き込みができるスコープを制限できる
		resultCh := make(chan int)
		go func() {
			defer close(resultCh)
			for i := 0; i < 3; i++ {
				resultCh <- i
			}
		}()
		return resultCh
	}

	// 受信操作専用のチャネルを受け取り、値を出力していく
	printer := func(resultCh <-chan int) {
		for v := range resultCh {
			fmt.Println(v)
		}
		fmt.Println("Done!!")
	}

	resultCh := chanOwner()
	printer(resultCh)
	// 実行結果
	// 0
	// 1
	// 2
	// Done!!

}
