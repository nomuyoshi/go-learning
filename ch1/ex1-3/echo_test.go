package main

import "testing"

var args = []string{"echo", "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

func BenchmarkLoopEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loopEcho(args)
	}
}

func BenchmarkJoinEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		joinEcho(args)
	}
}
