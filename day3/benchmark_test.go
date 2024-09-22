package main

import (
	"testing"
)

func BenchmarkA(b *testing.B) {

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		SolvePart1("../inputs/day3-1.txt")
	}
}

func BenchmarkB(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		SolvePart2("../inputs/day3-1.txt")
	}
}
func BenchmarkC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		SolvePart2Concurrent("../inputs/day3-1.txt")
	}
}
