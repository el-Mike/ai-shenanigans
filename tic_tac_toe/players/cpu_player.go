package players

import (
	"math/rand"
	"time"
)

const WIN_VALUE = 1
const DRAW_VALUE = 0
const LOSS_VALUE = -1

type ResultPair struct {
	result, move int
}

type ResultCache map[string]ResultPair

type CPUPlayer struct {
	sign         Sign
	board        Board
	cache        ResultCache
	stateChecker *StateChecker
}

func NewCPUPlayer(board Board, sign Sign) *CPUPlayer {
	return &CPUPlayer{
		sign:         sign,
		board:        board,
		cache:        ResultCache{},
		stateChecker: NewStateChecker(),
	}
}

func (cp *CPUPlayer) GetSign() Sign {
	return cp.sign
}

func (cp *CPUPlayer) Move() {
	board := cp.board.GetCopy()

	_, cell := cp.minmax(board, true)

	cp.board.PutSignByGridCell(cell, cp.sign)
}

func (cp *CPUPlayer) minmax(board Board, isMaximizer bool) (result int, move int) {
	hash := board.GetBoardHash()

	if cached, ok := cp.cache[hash]; ok {
		return cached.result, cached.move
	}

	gameState := cp.stateChecker.CheckState(board)

	if gameState == O_WON {
		return WIN_VALUE, -1
	}

	if gameState == X_WON {
		return LOSS_VALUE, -1
	}

	result, move = DRAW_VALUE, -1

	var targetResult int
	var currentSign Sign

	if isMaximizer {
		targetResult = WIN_VALUE
		currentSign = cp.sign
	} else {
		targetResult = LOSS_VALUE
		currentSign = X_SIGN
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

func (cp *CPUPlayer) RandomMove() {
	rand.Seed(time.Now().UnixNano())

	min := 1
	max := 9

	var cell int

	for {
		cell = rand.Intn(max-min*1) + min

		if cp.board.IsEmptyByGridCell(cell) {
			break
		}
	}

	cp.board.PutSignByGridCell(cell, cp.sign)
}
