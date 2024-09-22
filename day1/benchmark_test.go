package main

import (
	"testing"
)

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolvePart1("../inputs/day1-1.txt")
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolvePart2("../inputs/day1-1.txt")
	}
}
func BenchmarkC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2Replace("../inputs/day1-1.txt")
	}
}
