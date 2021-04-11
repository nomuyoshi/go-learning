package main

import "fmt"

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				fmt.Println("doWork")
				fmt.Println(s)
			}
		}()
		return completed
	}

	// doWorkにnilを渡しているのでstringsチャネルには絶対に文字列が書き込まれることはない
	// メインゴルーチンが生きている限り、doWork内で生成されたゴルーチンはメモリ内に残り続ける。
	// 今回のサンプルではメインゴルーチンはすぐに終了するが、実際のアプリケーションだと長期間稼働して
	// その間ゴルーチンを生成し続けるとじわじわメモリ使用率を高めていくことになってしまう。
	doWork(nil)
	fmt.Println("Done.")
}