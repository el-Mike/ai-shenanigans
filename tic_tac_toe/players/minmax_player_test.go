package players

import (
	"testing"

	"github.com/el-Mike/ai-shenanigans/tic_tac_toe/game"
)

func BenchmarkMinMax(b *testing.B) {
	board := game.NewBoard()
	cpu := NewMinmaxPlayer(board, game.O_SIGN, game.X_SIGN)

	for i := 0; i < b.N; i++ {
		cpu.minmax(board, true, -100, 100)
	}
}
