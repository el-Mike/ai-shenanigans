package main

type GameState string

const (
	NOT_FINISHED GameState = "NOT_FINISHED"
	X_WON        GameState = "X_WON"
	O_WON        GameState = "O_WON"
	DRAW         GameState = "DRAW"
)

type StateChecker struct{}

func NewStateChecker() *StateChecker {
	return &StateChecker{}
}

func (st *StateChecker) isWonRow(board Board, row int, sign Sign) bool {
	return board.IsSign(row, 0, sign) && board.IsSign(row, 1, sign) && board.IsSign(row, 2, sign)
}

func (st *StateChecker) isWonCol(board Board, col int, sign Sign) bool {
	return board.IsSign(0, col, sign) && board.IsSign(1, col, sign) && board.IsSign(2, col, sign)
}

func (st *StateChecker) isWonDiagFromLeft(board Board, sign Sign) bool {
	return board.IsSign(0, 0, sign) && board.IsSign(1, 1, sign) && board.IsSign(2, 2, sign)
}

func (st *StateChecker) isWonDiagFromRight(board Board, sign Sign) bool {
	return board.IsSign(0, 2, sign) && board.IsSign(1, 1, sign) && board.IsSign(2, 0, sign)
}

func (st *StateChecker) GetWinStateBySign(sign Sign) GameState {
	if sign == X_SIGN {
		return X_WON
	} else {
		return O_WON
	}
}

func (st *StateChecker) CheckState(board Board) GameState {
	signs := []Sign{X_SIGN, O_SIGN}

	for _, sign := range signs {
		for i := 0; i < BOARD_SIZE; i++ {
			if st.isWonRow(board, i, sign) || st.isWonCol(board, i, sign) {
				return st.GetWinStateBySign(sign)
			}
		}

		if st.isWonDiagFromLeft(board, sign) || st.isWonDiagFromRight(board, sign) {
			return st.GetWinStateBySign(sign)
		}
	}

	if !board.HasAnyEmpty() {
		return DRAW
	}

	return NOT_FINISHED
}

func (st *StateChecker) IsWon(board Board, sign Sign) bool {
	return st.CheckState(board) == st.GetWinStateBySign(sign)
}
