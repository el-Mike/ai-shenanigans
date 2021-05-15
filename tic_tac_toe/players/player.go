package players

import "github.com/el-Mike/ai-shenanigans/tic_tac_toe/board"

type Player interface {
	Move()
	GetSign() board.Sign
}
