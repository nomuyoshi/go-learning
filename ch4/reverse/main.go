package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // [5 4 3 2 1 0]

	b := []int{10, 9, 8, 7, 6, 5}
	reverse(b)
	fmt.Println(b) // [5 6 7 8 9 10]
}

// reverseは整数値のスライスを受け取り、要素を逆順にする
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
