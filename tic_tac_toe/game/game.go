package game

import (
	"fmt"
)

type Game struct {
	board        Board
	stateChecker *StateChecker
	renderer     Renderer
}

func NewGame(board Board, stateChecker *StateChecker, renderer Renderer) *Game {
	return &Game{
		board:        board,
		stateChecker: stateChecker,
		renderer:     renderer,
	}
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

	g.renderer.Render(g.board, currentSign)

	for {
		currentPlayer.Move()

		g.renderer.Render(g.board, currentSign)

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
