package main

import (
	"fmt"
	"time"
)

// 値を受信する処理をdoneチャネルを使ってキャンセルする例
func main() {
	// doneチャネルは処理終了を伝えるためのチャネル。（＝キャンセルを伝える）
	print := func(
		done <-chan interface{}, // 受信専用チャネル
		strings <-chan string, // 受信専用チャネル
	) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("print exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					// stringsチャネルから受信した文字列を出力
					fmt.Println(s)
				case <-done:
					// doneチャネルが閉じられたら抜ける
					fmt.Println("close done")
					return
				}
			}
		}()
		return terminated
	}

	strCh := make(chan string)
	done := make(chan interface{})
	terminated := print(done, strCh)

	for _, s := range []string{"golang", "ruby", "python"} {
		strCh <- s
	}
	// printをキャンセルするゴルーチン
	// ちなみに、このゴルーチンを <-terminatedの後に生成するとプログラムはDeadlockしてしまう。
	// → terminatedに値が送られる、閉じられることがないから。
	go func() {
		// 1秒待って、doneチャネルを閉じる
		// doneチャネルを閉じることで、printに処理の終了（キャンセル）を伝える
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling print goroutine...")
		close(done)
	}()

	// terminatedチャネルが閉じられるまで待機
	<-terminated
	fmt.Println("Done.")
}
