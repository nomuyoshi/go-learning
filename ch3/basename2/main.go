package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// bufio パッケージは，I/O で役に立つ機能を提供してくれる
	// ここでは、標準入力から1行ずつ読み込むScannerを生成している
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
}

// basenameはディレクトリ要素と接尾辞を取り除く
// "dir1/dir2/filename.ext" => filename
func basename(s string) string {
	// 引数sの後ろから探索して、'/'があったら、その'/'より後ろの文字のみにする
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// 後ろから探索して、'.'があったら、'.'より前の文字のみにする
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
