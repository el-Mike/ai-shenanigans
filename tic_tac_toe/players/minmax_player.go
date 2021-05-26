package players

import (
	"github.com/el-Mike/ai-shenanigans/tic_tac_toe/game"
)

const WIN_VALUE = 1
const DRAW_VALUE = 0
const LOSS_VALUE = -1

const NOOP_ACTION = -1

// Initials to Alpha and Beta need to be the worst cases for bot max and min
// players, respectively - therefore for maximizing player tends to -Infinity, and for
// minimalizing player - to Infinity.
const INITIAL_ALPHA = -2
const INITIAL_BETA = 2

type ResultPair struct {
	result, move int
}

type ResultCache map[string]ResultPair

type MinmaxPlayer struct {
	sign         game.Sign
	opponentSign game.Sign
	board        game.Board
	cache        ResultCache
	stateChecker *game.StateChecker
}

func NewMinmaxPlayer(board game.Board, sign game.Sign, opponentSign game.Sign) *MinmaxPlayer {
	return &MinmaxPlayer{
		sign:         sign,
		opponentSign: opponentSign,
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

	_, cell := cp.minmax(board, true, INITIAL_ALPHA, INITIAL_BETA)

	cp.board.PutSignByGridCell(cell, cp.sign)
}

func (cp *MinmaxPlayer) minmax(board game.Board, isMaximizer bool, alpha int, beta int) (result int, move int) {
	hash := board.GetBoardHash()

	if cached, ok := cp.cache[hash]; ok {
		return cached.result, cached.move
	}

	gameState := cp.stateChecker.CheckState(board)

	if gameState == cp.stateChecker.GetWinStateBySign(cp.sign) {
		return WIN_VALUE, NOOP_ACTION
	}

	if gameState == cp.stateChecker.GetWinStateBySign(cp.opponentSign) {
		return LOSS_VALUE, NOOP_ACTION
	}

	result, move = DRAW_VALUE, NOOP_ACTION

	var targetResult int
	var currentSign game.Sign

	if isMaximizer {
		targetResult = WIN_VALUE
		currentSign = cp.sign
	} else {
		targetResult = LOSS_VALUE
		currentSign = cp.opponentSign
	}

	availableCells := board.GetEmptyCells()

	for i := 0; i < len(availableCells); i++ {
		targetCell := availableCells[i]

		newBoard := board.GetCopy()
		newBoard.PutSignByGridCell(targetCell, currentSign)

		currentResult, _ := cp.minmax(newBoard, !isMaximizer, alpha, beta)

		if isMaximizer {
			alpha = currentResult
		} else {
			beta = currentResult
		}

		if (isMaximizer && (currentResult > result)) ||
			(!isMaximizer && (currentResult < result)) ||
			move == NOOP_ACTION {
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
