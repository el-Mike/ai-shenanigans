package main

func main() {
	board := NewBoard()
	stateChecker := NewStateChecker()

	game := &Game{
		board:        board,
		stateChecker: stateChecker,
	}

	game.Start()
}
