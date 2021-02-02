package main

import (
	"fmt"
	"sort"
)

func main() {
	var names []string
	names = []string{"Ryo", "Akari", "Chen", "Eri", "Gon", "Jin", "Ken"}
	fmt.Printf("ソート前: %s\n", names) // ソート前: [Ryo Akari Chen Eri Gon Jin Ken]
	sort.Strings(names)             // sort.Sort(sort.StringSlice(names))
	fmt.Printf("ソート後: %s\n", names) // ソート後: [Akari Chen Eri Gon Jin Ken Ryo]

	// 逆順にソート
	sort.Sort(sort.Reverse(sort.StringSlice(names)))
	fmt.Printf("逆順にソート: %s\n", names) // 逆順にソート: [Ryo Ken Jin Gon Eri Chen Akari]
}
