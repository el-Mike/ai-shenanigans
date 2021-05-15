package game

import (
	"fmt"

	"github.com/el-Mike/ai-shenanigans/tic_tac_toe/board"
	"github.com/el-Mike/ai-shenanigans/tic_tac_toe/players"
)

type Game struct {
	board        board.Board
	stateChecker *StateChecker
}

func (g *Game) renderBoard(currentSign board.Sign) {
	fmt.Print(board.RENDER_SEPARATOR)
	fmt.Printf("Current sign: %v\n\n", currentSign)

	for _, row := range g.board {
		fmt.Print("    ")
		for j, sign := range row {
			if sign == "" {
				fmt.Print(board.EMPTY_SIGN)
			} else {
				fmt.Print(sign)
			}

			if j < (board.BOARD_SIZE - 1) {
				fmt.Print(" | ")
			} else {
				fmt.Print("\n")
			}
		}
	}

	fmt.Print(board.RENDER_SEPARATOR)
}

func (g *Game) Move(cell int, sign board.Sign) {
	g.board.PutSignByGridCell(cell, sign)
}

// @TODO:
// 1. Add randomness to MinMax.
// 2. Add alpha-beta pruning.
func (g *Game) Start() {
	player := players.NewHumanPlayer(g.board, board.X_SIGN)
	cpuPlayer := players.NewCPUPlayer(g.board, board.O_SIGN)

	var currentPlayer players.Player
	currentPlayer = player

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

		if _, ok := currentPlayer.(*players.HumanPlayer); ok {
			currentPlayer = cpuPlayer
			currentSign = currentPlayer.GetSign()
		} else {
			currentPlayer = player
			currentSign = currentPlayer.GetSign()
		}
	}
}
