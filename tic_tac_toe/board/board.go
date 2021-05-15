package board

import "fmt"

const BOARD_SIZE = 3

const RENDER_SEPARATOR = "\n================\n"

type Sign string

const (
	X_SIGN     Sign = "X"
	O_SIGN     Sign = "O"
	EMPTY_SIGN Sign = "_"
)

type Row []Sign
type Board []Row

func NewBoard() Board {
	return Board{
		Row{EMPTY_SIGN, EMPTY_SIGN, EMPTY_SIGN},
		Row{EMPTY_SIGN, EMPTY_SIGN, EMPTY_SIGN},
		Row{EMPTY_SIGN, EMPTY_SIGN, EMPTY_SIGN},
	}
}

func (b Board) forEvery(cb func(row, col int)) {
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			cb(i, j)
		}
	}
}

func (b Board) CellToCoords(cell int) (row int, col int) {
	row = (cell - 1) / 3
	col = (cell - 1) % 3

	return row, col
}

func (b Board) CoordsToCell(row, col int) int {
	return (row * 3) + (col + 1)
}

func (b Board) GetSign(row, col int) Sign {
	return b[row][col]
}

func (b Board) IsSign(row, col int, sign Sign) bool {
	return b[row][col] == sign
}

func (b Board) IsSignByGridCell(cell int, sign Sign) bool {
	row, col := b.CellToCoords(cell)

	return b.IsSign(row, col, sign)
}

func (b Board) IsEmpty(row, col int) bool {
	return b[row][col] == EMPTY_SIGN
}

func (b Board) IsEmptyByGridCell(cell int) bool {
	row, col := b.CellToCoords(cell)

	return b.IsEmpty(row, col)
}

func (b Board) putSign(row, col int, sign Sign) {
	b[row][col] = sign
}

func (b Board) PutSignByGridCell(cell int, sign Sign) {
	row, col := b.CellToCoords(cell)

	b.putSign(row, col, sign)
}

func (b Board) HasAnyEmpty() bool {
	hasEmpty := false

	b.forEvery(func(row, col int) {
		if b.IsEmpty(row, col) {
			hasEmpty = true
		}
	})

	return hasEmpty
}

func (b Board) GetEmptyCells() []int {
	var cells []int

	b.forEvery(func(row, col int) {
		if b.IsEmpty(row, col) {
			cells = append(cells, b.CoordsToCell(row, col))
		}
	})

	return cells
}

func (b Board) GetCopy() Board {
	newBoard := NewBoard()

	b.forEvery(func(row, col int) {
		newBoard[row][col] = b[row][col]
	})

	return newBoard
}

func (b Board) GetBoardHash() string {
	hash := ""

	b.forEvery(func(row, col int) {
		sign := b.GetSign(row, col)
		hash += fmt.Sprintf("%v_%v_%v", row, col, sign)
	})

	return hash
}
