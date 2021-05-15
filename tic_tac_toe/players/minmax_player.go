package players

import (
	"github.com/el-Mike/ai-shenanigans/tic_tac_toe/game"
)

const WIN_VALUE = 1
const DRAW_VALUE = 0
const LOSS_VALUE = -1

type ResultPair struct {
	result, move int
}

type ResultCache map[string]ResultPair

type MinmaxPlayer struct {
	sign         game.Sign
	board        game.Board
	cache        ResultCache
	stateChecker *game.StateChecker
}

func NewMinmaxPlayer(board game.Board, sign game.Sign) *MinmaxPlayer {
	return &MinmaxPlayer{
		sign:         sign,
		board:        board,
		cache:        ResultCache{},
		stateChecker: game.NewStateChecker(),
	}
}

func (cp *MinmaxPlayer) GetSign() game.Sign {
	return cp.sign
}

func (cp *MinmaxPlayer) Move() {
	board := cp.board.GetCopy()

	_, cell := cp.minmax(board, true)

	cp.board.PutSignByGridCell(cell, cp.sign)
}

func (cp *MinmaxPlayer) minmax(board game.Board, isMaximizer bool) (result int, move int) {
	hash := board.GetBoardHash()

	if cached, ok := cp.cache[hash]; ok {
		return cached.result, cached.move
	}

	gameState := cp.stateChecker.CheckState(board)

	if gameState == game.O_WON {
		return WIN_VALUE, -1
	}

	if gameState == game.X_WON {
		return LOSS_VALUE, -1
	}

	result, move = DRAW_VALUE, -1

	var targetResult int
	var currentSign game.Sign

	if isMaximizer {
		targetResult = WIN_VALUE
		currentSign = cp.sign
	} else {
		targetResult = LOSS_VALUE
		currentSign = game.X_SIGN
	}

	availableCells := board.GetEmptyCells()

	for i := 0; i < len(availableCells); i++ {
		targetCell := availableCells[i]

		newBoard := board.GetCopy()
		newBoard.PutSignByGridCell(targetCell, currentSign)

		currentResult, _ := cp.minmax(newBoard, !isMaximizer)

		if (isMaximizer && (currentResult > result)) ||
			(!isMaximizer && (currentResult < result)) ||
			move == -1 {
			result, move = currentResult, targetCell
		}

		if result == targetResult {
			cp.cache[hash] = ResultPair{result, move}

			return result, move
		}

		cp.cache[hash] = ResultPair{result, move}
	}

	return result, move
}
