package main

import "testing"

func BenchmarkDay3Part2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		process()
	}
}
