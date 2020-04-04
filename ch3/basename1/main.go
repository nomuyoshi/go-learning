package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	slash := strings.LastIndex(s, "/") // => 見つからなければ-1が返る
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
