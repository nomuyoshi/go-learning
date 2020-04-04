package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(comma(os.Args[i]))
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	// 最後の3文字以外の部分文字列をcomma関数にわたして再帰的に呼び出している
	return comma(s[:n-3]) + "," + s[n-3:]
}
