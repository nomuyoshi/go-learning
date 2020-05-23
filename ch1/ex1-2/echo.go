package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("index: %d, value: %s\n", i, arg)
	}
}
