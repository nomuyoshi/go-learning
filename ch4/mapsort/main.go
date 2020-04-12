package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"ziro":  11,
		"mike":  30,
		"alice": 18,
		"taro":  24,
	}
	// mapのループは繰り返し順は定義されておらず、ランダム（実装ごとに違った順序になる）
	// そのため、決まった順序で繰り返したい場合は次のようにキーをsortしたスライスを用意する

	// 必要な容量はわかっているので、最初から必要な大きさのスライスを定義した方が効率的
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
