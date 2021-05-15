package main

import (
	"github.com/el-Mike/ai-shenanigans/tic_tac_toe/game"
	"github.com/el-Mike/ai-shenanigans/tic_tac_toe/players"
)

func main() {
	board := game.NewBoard()
	stateChecker := game.NewStateChecker()

	g := game.NewGame(board, stateChecker)

	playerOne := players.NewHumanPlayer(board, game.X_SIGN)
	playerTwo := players.NewMinmaxPlayer(board, game.O_SIGN)

	g.Start(playerOne, playerTwo)
}
