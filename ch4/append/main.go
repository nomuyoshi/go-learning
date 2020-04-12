package main

import "fmt"

func main() {
	var x []int
	x = appendInt(x, 0, 1, 2, 3, 4, 5)
	fmt.Printf("cap=%d len=%d %v\n", cap(x), len(x), x)
}

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// 容量に余裕がある場合は、スライスを拡張する
		z = x[:zlen]
	} else {
		// 十分な容量がない場合は、新たな配列を割り当てる。
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
