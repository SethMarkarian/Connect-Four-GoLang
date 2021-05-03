package main

import "fmt"

type Game struct {
	board      [][]string
	playerTurn int
}

func createGame() Game {
	newGame := Game{}
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

func printBoard(game Game) {
	for i := 0; i < 6; i++ {
		fmt.Println(game.board[i])
	}
}

func main() {
	g := createGame()
	printBoard(g)
}
