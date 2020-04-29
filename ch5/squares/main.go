package main

import "fmt"

func main() {
	f := squares()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 4
	fmt.Println(f()) // 9
	fmt.Println(f()) // 16
	fmt.Println(f()) // 25
}

// クロージャー
// squaresは無名関数を返す。「func() int」が返り値の型
// 無名関数はsquares関数内のローカル変数にアクセス、更新可能で
// 呼び出されるごとにローカル変数xを加算する（状態をもつことができる）
// ローカル変数xはsquaresが戻ったあとも、main内に存在する。
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
