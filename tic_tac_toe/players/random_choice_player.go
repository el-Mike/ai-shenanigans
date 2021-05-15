package players

import (
	"math/rand"
	"time"

	"github.com/el-Mike/ai-shenanigans/tic_tac_toe/game"
)

type RandomChoicePlayer struct {
	board game.Board
	sign  game.Sign
}

func NewRandomChicePlayer(board game.Board, sign game.Sign) *RandomChoicePlayer {
	return &RandomChoicePlayer{
		board: board,
		sign:  sign,
	}
}

func (rcp *RandomChoicePlayer) GetSign() game.Sign {
	return rcp.sign
}

func (rcp *RandomChoicePlayer) Move() {
	rand.Seed(time.Now().UnixNano())

	min := 1
	max := 9

	var cell int

	for i := 0; ; i++ {
		cell = rand.Intn(max-min*1) + min

		if rcp.board.IsEmptyByGridCell(cell) {
			break
		}

		// Limit number of iterations to 20 - if no empty cell was found,
		// get first empty cell and put the sign inside it.
		// It prevents the loop for taking too long.
		if i == 20 {
			cell = rcp.board.GetFirstEmpty()
			break
		}
	}

	rcp.board.PutSignByGridCell(cell, rcp.sign)
}
