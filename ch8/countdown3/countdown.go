package main

import (
	"fmt"
	"os"
	"time"
)

// NOTE: カウントダウンが終了してもTick関数が生成しているゴルーチンは生存している
// つまり、受信するゴルーチンのないtickチャネルに無駄に送信を試みている -> ゴルーチンのリーク
// これを避けるためには、次のようにする
// ticker := time.NewTicker(1 * time.Second)
// <-ticker.C // tickerのチャネルから受信
// ticker.Stop() // 不要になったらtickerのゴルーチンを終了させる

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort.")
	// time.Tickはチャネルを返し、定期的にイベントを送信する
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		// tickからイベントを受信する
		// tickには1秒ごとにイベントが送信されているので、1秒間隔でPrintされることになる
		fmt.Println(countdown)
		select {
		case <-tick:
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
