package renderers

import (
	"fmt"

	"github.com/el-Mike/ai-shenanigans/tic_tac_toe/game"
)

const RENDER_SEPARATOR = "\n================\n"

type ConsoleRenderer struct{}

func NewConsoleRenderer() *ConsoleRenderer {
	return &ConsoleRenderer{}
}

func (cr *ConsoleRenderer) Render(board game.Board, currentSign game.Sign) {
	fmt.Print(RENDER_SEPARATOR)
	fmt.Printf("Current sign: %v\n\n", currentSign)

	for _, row := range board {
		fmt.Print("    ")
		for j, sign := range row {
			if sign == "" {
				fmt.Print(game.EMPTY_SIGN)
			} else {
				fmt.Print(sign)
			}

			if j < (game.BOARD_SIZE - 1) {
				fmt.Print(" | ")
			} else {
				fmt.Print("\n")
			}
		}
	}

	fmt.Print(RENDER_SEPARATOR)
}
