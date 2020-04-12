package main

import "fmt"

func main() {
	data1 := []string{"ruby", "", "php", "", "python", "golang", ""}
	fmt.Println("----- nonempty -----")
	fmt.Printf("nonempty: %q\n", nonempty(data1))
	fmt.Printf("original: %q\n", data1)

	data2 := []string{"ruby", "", "php", "", "python", "golang", ""}
	fmt.Println("----- nonempty2 -----")
	fmt.Printf("nonempty2: %q\n", nonempty2(data2))
	fmt.Printf("original: %q\n", data2)

	data3 := []string{"ruby", "", "php", "", "python", "golang", ""}
	fmt.Println("----- nonempty3 -----")
	fmt.Printf("nonempty3: %q\n", nonempty3(data3))
	fmt.Printf("original: %q\n", data3)
}

// nonempty は文字列のリストを受け取り、空文字列を含まないリストを返す
// 入力スライスと出力スライスは同じ基底配列を参照している実装となっている。
// そのため、入力スライスも部分的に上書きされる。
// 次のようにnonemptyに渡されたスライス変数に代入するように書く
// data = nonempty(data)
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}

func nonempty2(strings []string) []string {
	// 長さ0、容量はもとのスライスと同じスライスのコピーを用意
	// ただし、同じ基底配列を参照
	out := strings[:0]
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(out), cap(out), out)
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}

	return out
}

func nonempty3(strings []string) []string {
	// 長さ0、容量はもとのスライスと同じ新しいスライスをmake関数を使って宣言
	// 基底配列も異なるため、入力スライスが上書かれることはない
	out := make([]string, 0, cap(strings))
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(out), cap(out), out)
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}

	return out
}
