package players

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/el-Mike/ai-shenanigans/tic_tac_toe/game"
)

type HumanPlayer struct {
	sign  game.Sign
	board game.Board
}

func NewHumanPlayer(board game.Board, sign game.Sign) *HumanPlayer {
	return &HumanPlayer{
		sign:  sign,
		board: board,
	}
}

func (p *HumanPlayer) GetSign() game.Sign {
	return p.sign
}

func (p *HumanPlayer) Move() {
	cell := p.readCell()

	p.board.PutSignByGridCell(cell, p.sign)
}

func (p *HumanPlayer) readCell() int {
	scanner := bufio.NewScanner(os.Stdin)

	var cell int
	var err error

	for {
		fmt.Print("Enter grid cell (1 - 9): ")

		scanner.Scan()
		cellStr := scanner.Text()

		cell, err = strconv.Atoi(cellStr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if cell < 1 || cell > 9 {
			fmt.Println("Wrong value, try again.")
			continue
		} else if !p.board.IsEmptyByGridCell(cell) {
			fmt.Println("Cell already taken, try again.")
			continue
		} else {
			break
		}
	}

	return cell
}
