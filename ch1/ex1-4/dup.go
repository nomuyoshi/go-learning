package main

import (
	"bufio"
	"fmt"
	"os"
)

// 標準入力orファイルから2回以上あらわれる行を出現回数とともに表示する
func main() {
	counts := make(map[string]int)
	// key: 重複行、value: ファイル名
	foundIn := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		// ファイル指定がなければ、標準入力の内容をカウントする
		countLines(os.Stdin, counts, foundIn)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts, foundIn)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, foundIn[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, foundIn map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if _, ok := foundIn[line]; !ok {
			foundIn[line] = f.Name()
		}
	}
}
