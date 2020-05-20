package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	// time.Tickはチャネルを返し、定期的にイベントを送信する
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		// tickからイベントを受信する
		// tickには1秒ごとにイベントが送信されているので、1秒間隔でPrintされることになる
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
