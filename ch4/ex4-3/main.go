package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a) // [5 4 3 2 1]
}

func reverse(ary *[5]int) {
	for i, j := 0, 4; i < j; i, j = i+1, j-1 {
		ary[i], ary[j] = ary[j], ary[i]
	}
}
