package main

import (
	"testing"
)

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolvePart1("../inputs/day2-2.txt")
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolvePart2("../inputs/day2-2.txt")
	}
}
