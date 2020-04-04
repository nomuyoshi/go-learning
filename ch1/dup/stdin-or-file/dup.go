package main

import (
	"bufio"
	"fmt"
	"os"
)

// 標準入力orファイルから2回以上あらわれる行を出現回数とともに表示する
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		// ファイル指定がなければ、標準入力の内容をカウントする
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
