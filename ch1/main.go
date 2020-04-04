package main

import (
	"fmt"
)

func main() {
	// counts := make(map[string]int)
	counts := map[string]int{"america": 1}

	updateCounts(counts)
	for key, v := range counts {
		fmt.Printf("%s: %d\n", key, v)
	}
}

func updateCounts(counts map[string]int) {
	counts["test"] = 2
	counts["japan"] = 1000
}
