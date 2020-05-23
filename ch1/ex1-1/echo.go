package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := strings.Join(os.Args[1:], " ")
	fmt.Printf("%s: %s\n", os.Args[0], args)
}
