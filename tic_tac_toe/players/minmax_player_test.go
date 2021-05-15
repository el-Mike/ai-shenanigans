package players

import (
	"testing"

	"github.com/el-Mike/ai-shenanigans/tic_tac_toe/board"
)

func BenchmarkMinMax(b *testing.B) {
	board := board.NewBoard()
	cpu := NewCPUPlayer(board, board.O_SIGN)

	for i := 0; i < b.N; i++ {
		cpu.minmax(board, true)
	}
}
