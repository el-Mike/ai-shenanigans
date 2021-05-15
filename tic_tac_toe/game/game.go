package game

import "fmt"

type Game struct {
	board        Board
	stateChecker *StateChecker
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
func (g *Game) Start() {
	player := NewHumanPlayer(g.board, X_SIGN)
	cpuPlayer := NewCPUPlayer(g.board, O_SIGN)

	var currentPlayer Player
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

		if _, ok := currentPlayer.(*HumanPlayer); ok {
			currentPlayer = cpuPlayer
			currentSign = currentPlayer.GetSign()
		} else {
			currentPlayer = player
			currentSign = currentPlayer.GetSign()
		}
	}
}
