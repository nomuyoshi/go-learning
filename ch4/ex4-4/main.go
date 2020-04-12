package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5}
	slice = rotate(slice, 3)
	fmt.Println(slice) // [3 4 5 0 1 2]
}

// rotate は受け取ったインデックスの値が先頭になるように
// 要素順を変更したスライスを返す
func rotate(slice []int, i int) []int {
	return append(slice[i:], slice[:i]...)
}
