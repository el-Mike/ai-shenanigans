package game

import (
	"fmt"
)

type Game struct {
	board        Board
	stateChecker *StateChecker
}

func NewGame(board Board, stateChecker *StateChecker) *Game {
	return &Game{
		board:        board,
		stateChecker: stateChecker,
	}
}

func (g *Game) renderBoard(currentSign Sign) {
	fmt.Print(RENDER_SEPARATOR)
	fmt.Printf("Current sign: %v\n\n", currentSign)

	for _, row := range g.board {
		fmt.Print("    ")
		for j, sign := range row {
			if sign == "" {
				fmt.Print(EMPTY_SIGN)
			} else {
				fmt.Print(sign)
			}

			if j < (BOARD_SIZE - 1) {
				fmt.Print(" | ")
			} else {
				fmt.Print("\n")
			}
		}
	}

	fmt.Print(RENDER_SEPARATOR)
}

func (g *Game) Move(cell int, sign Sign) {
	g.board.PutSignByGridCell(cell, sign)
}

// @TODO:
// 1. Add randomness to MinMax.
// 2. Add alpha-beta pruning.
func (g *Game) Start(playerOne, playerTwo Player) {
	var currentPlayer Player
	currentPlayer = playerOne

	currentSign := currentPlayer.GetSign()

	g.renderBoard(currentSign)

	for {
		currentPlayer.Move()

		g.renderBoard(currentSign)

		if g.stateChecker.IsWon(g.board, currentSign) {
			fmt.Printf("%v won!\n", currentSign)
			break
		}

		if !g.board.HasAnyEmpty() {
			fmt.Println("Draw!")
			break
		}

		if currentPlayer == playerOne {
			currentPlayer = playerTwo
		} else {
			currentPlayer = playerOne
		}

		currentSign = currentPlayer.GetSign()
	}
}
