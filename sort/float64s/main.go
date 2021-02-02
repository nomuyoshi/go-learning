package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []float64{1.1, 1.4, 1.2, 1.5, 1.3}
	fmt.Printf("ソート前: %g\n", numbers) // ソート前: [1.1 1.4 1.2 1.5 1.3]
	sort.Float64s(numbers)            // sort.Sort(sort.Float64Slice(numbers))
	fmt.Printf("ソート後: %g\n", numbers) // ソート後: ソート後: [1.1 1.2 1.3 1.4 1.5]

	// 逆順にソート
	sort.Sort(sort.Reverse(sort.Float64Slice(numbers)))
	fmt.Printf("逆順にソート: %g\n", numbers) // 逆順にソート: [1.5 1.4 1.3 1.2 1.1]
}
