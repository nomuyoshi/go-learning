package main

import "fmt"

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	// ch1, ch2ともに宣言しただけで、何も送信されてこない（受信することはない）
	// そのため、default節が実行される
	select {
	case <-ch1:
		fmt.Println("ch1")
	case <-ch2:
		fmt.Println("ch2")
	default:
		fmt.Println("default")
	}
}
