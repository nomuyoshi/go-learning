package main

import "fmt"

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	var count1, count2 int
	close(ch1)
	close(ch2)
	for i := 0; i < 100; i++ {
		// ch1, ch2ともに閉じられているので、準備が整っている（すぐに受信できる）
		select {
		case <-ch1:
			count1++
		case <-ch2:
			count2++
		}
	}

	fmt.Printf("count1 = %d, count2 = %d\n", count1, count2)
}
