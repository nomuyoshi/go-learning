package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{4, 1, 2, 5, 3}
	fmt.Printf("ソート前: %d\n", numbers) // ソート前: [4 1 2 5 3]
	sort.Ints(numbers)                // sort.Sort(sort.IntSlice(numbers))
	fmt.Printf("ソート後: %d\n", numbers) // ソート後: ソート後: [1 2 3 4 5]

	// 逆順にソート
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Printf("逆順にソート: %d\n", numbers) // 逆順にソート: [5 4 3 2 1]
}
