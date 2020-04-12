package main

import "fmt"

func main() {
	strings := []string{"ruby", "php", "php", "php", "python", "python", "ruby", "go", "go"}
	fmt.Println(uniq(strings))
}

// uniq は隣接している重複を除去する
func uniq(slice []string) []string {
	out := slice[:0]
	for i, s := range slice {
		if i == 0 {
			out = append(out, s)
			continue
		}

		if slice[i-1] != s {
			out = append(out, s)
		}
	}

	return out
}
