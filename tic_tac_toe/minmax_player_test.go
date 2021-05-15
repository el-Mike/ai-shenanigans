package main

import "testing"

func BenchmarkMinMax(b *testing.B) {
	board := NewBoard()
	cpu := NewCPUPlayer(board, O_SIGN)

	for i := 0; i < b.N; i++ {
		cpu.minmax(board, true)
	}
}
