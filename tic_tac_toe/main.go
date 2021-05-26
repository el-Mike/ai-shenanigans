package main

import (
	"github.com/el-Mike/ai-shenanigans/tic_tac_toe/game"
	"github.com/el-Mike/ai-shenanigans/tic_tac_toe/players"
	"github.com/el-Mike/ai-shenanigans/tic_tac_toe/renderers"
)

func main() {
	board := game.NewBoard()
	stateChecker := game.NewStateChecker()
	renderer := renderers.NewConsoleRenderer()

	g := game.NewGame(board, stateChecker, renderer)

	playerOne := players.NewHumanPlayer(board, game.X_SIGN)
	playerTwo := players.NewMinmaxPlayer(board, game.O_SIGN, playerOne.GetSign())

	g.Start(playerOne, playerTwo)
}
