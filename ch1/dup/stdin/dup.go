package main

import (
	"bufio"
	"fmt"
	"os"
)

// 標準入力から2回以上あらわれる行を出現回数とともに表示する
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// Scan() を呼び出すごとに1行ずつ読み込み、行末から改行文字を取り除く。
	// 行があればtrue、なければfalseを返すので無限ループにはならない。
	for input.Scan() {
		// intのゼロ値は0なので、mapにキーがなくても問題ない
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
