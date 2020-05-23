package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("loop echo")
	fmt.Println(loopEcho(os.Args[1:]))
	fmt.Println("join echo")
	fmt.Println(joinEcho(os.Args[1:]))
}

func joinEcho(args []string) string {
	return strings.Join(args, " ")
}

func loopEcho(args []string) string {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}

	return s
}
