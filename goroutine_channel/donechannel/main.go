package main

import (
	"fmt"
	"time"
)

func main() {
	// doneチャネルは処理終了を伝えるためのチャネル。（＝キャンセルを伝える）
	doWork := func(
		done <-chan interface{}, // 受信専用チャネル
		strings <-chan string, // 受信専用チャネル
	) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					// stringsチャネルから受信した文字列を出力
					fmt.Println(s)
				case <-done:
					fmt.Println("close done")
					// doneチャネルが閉じられたら抜ける
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)
	go func() {
		// 1秒待って、doneチャネルを閉じる
		// doneチャネルを閉じることで、doWorkに処理の終了（キャンセル）を伝える
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	// terminatedチャネルが閉じられるまで待機
	<-terminated
	fmt.Println("Done.")
}
