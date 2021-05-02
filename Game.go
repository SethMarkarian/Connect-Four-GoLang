package main

type Game struct {
	board      [][]string
	playerTurn int
}

func createGame() Game {
	var newGame = Game{}
	newGame.board = [][]string{
		{"_", "_", "_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_", "_", "_"},
	}
	newGame.playerTurn = 1
	return newGame
}
