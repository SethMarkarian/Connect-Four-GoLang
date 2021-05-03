package main

import "fmt"

type Game struct {
	board      [][]string
	playerTurn int
}

func createGame() Game {
	newGame := Game{}
	newGame.board = [][]string{
		{"_", "_", "_", "_", "_", "_", "_"}, {"_", "_", "_", "_", "_", "_", "_"}, {"_", "_", "_", "_", "_", "_", "_"}, {"_", "_", "_", "_", "_", "_", "_"}, {"_", "_", "_", "_", "_", "_", "_"}, {"_", "_", "_", "_", "_", "_", "_"},
	}
	newGame.playerTurn = 1
	return newGame
}

func printBoard(game Game) {
	for i := 0; i < 6; i++ {
		fmt.Println(game.board[i])
	}
}

func addPiece(index int, piece string, game Game) Game {
	for i := 5; i >= 0; i-- {
		if game.board[i][index] == "_" {
			game.board[i][index] = piece
			break
		}
	}
	return game
}

func main() {
	g := createGame()
	g = addPiece(2, "x", g)
	g = addPiece(3, "o", g)
	g = addPiece(2, "x", g)
	printBoard(g)
}
