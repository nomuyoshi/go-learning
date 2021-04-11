package main

import (
	"fmt"
	"time"
)

func main() {
	idCh := make(chan int)
	ids := []int{1, 2, 3, 4, 5}
	go func() {
		// 1秒ごとにチャネルにidを送信
		// すべて送信したら閉じる
		defer close(idCh)
		for _, id := range ids {
			time.Sleep(1 * time.Second)
			idCh <- id
		}
	}()

loop:
	for {
		select {
		case id, ok := <-idCh:
			// チャネルが閉じられたら終了
			if !ok {
				break loop
			}
			fmt.Println(id)
		default:
			// 何もしない。ブロックしないためにdefault節を用意しておく
		}
	}

	fmt.Println("Do something...")
	fmt.Println("Done!!")
}
