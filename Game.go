package main

import (
	"fmt"
)

type Game struct {
	board          [][]string
	piecesInColumn []int
	done           bool
}

func createGame() Game {
	newGame := Game{}
	newGame.board = [][]string{
		{"_", "_", "_", "_", "_", "_", "_"}, {"_", "_", "_", "_", "_", "_", "_"}, {"_", "_", "_", "_", "_", "_", "_"}, {"_", "_", "_", "_", "_", "_", "_"}, {"_", "_", "_", "_", "_", "_", "_"}, {"_", "_", "_", "_", "_", "_", "_"},
	}
	newGame.done = false
	newGame.piecesInColumn = []int{0, 0, 0, 0, 0, 0, 0}
	return newGame
}

func printBoard(game Game) {
	for i := 0; i < 6; i++ {
		fmt.Println(game.board[i])
	}
	fmt.Println(" 1 2 3 4 5 6 7")
}

func boardFilled(game Game) bool {
	isFull := true
	for i := 0; i < 7; i++ {
		if game.piecesInColumn[i] <= 5 {
			isFull = false
			break
		}
	}
	return isFull
}

func addPiece(index int, piece string, game Game) Game {
	for i := 5; i >= 0; i-- {
		if game.board[i][index] == "_" {
			game.board[i][index] = piece
			break
		}
	}
	game.piecesInColumn[index]++
	return game
}

func gameWin(game Game) int {
	//horizontal check
	for _, line := range game.board {
		for i := 0; i <= 7-4; i++ {
			if compare(line[i:i+4], "X") {
				game.done = true
				return 1
			} else if compare(line[i:i+4], "O") {
				game.done = true
				return 2
			}
		}
	}

	//vertical check
	for i := 0; i < 7; i++ {
		currentRow := getVerticalLine(game, i)
		for j := 0; j <= 7-4; j++ {
			if compare(currentRow[j:j+4], "X") {
				game.done = true
				return 1
			} else if compare(currentRow[j:j+4], "O") {
				game.done = true
				return 2
			}
		}
	}

	//diagonal L-R check
	for i := 5; i >= 3; i-- {
		for j := 0; j <= 3; j++ {
			currentDiag := getDiagLR(game, i, j)
			if compare(currentDiag, "X") {
				game.done = true
				return 1
			} else if compare(currentDiag, "O") {
				game.done = true
				return 2
			}
		}
	}

	//diagonal R-L check
	for i := 5; i >= 3; i-- {
		for j := 6; j >= 3; j-- {
			currentDiag := getDiagRL(game, i, j)
			if compare(currentDiag, "X") {
				game.done = true
				return 1
			} else if compare(currentDiag, "O") {
				game.done = true
				return 2
			}
		}
	}
	return -1
}

func compare(slice []string, str string) bool {
	if slice == nil {
		return false
	}
	for _, char := range slice {
		if char != str {
			return false
		}
	}
	return true
}

func getVerticalLine(game Game, col int) []string {
	returnSlice := []string{}
	for i := 0; i < 6; i++ {
		returnSlice = append(returnSlice, game.board[i][col])
	}
	return returnSlice
}

func getDiagLR(game Game, x int, y int) []string {
	returnSlice := []string{}
	returnSlice = append(returnSlice, game.board[x][y], game.board[x-1][y+1], game.board[x-2][y+2], game.board[x-3][y+3])
	return returnSlice
}

func getDiagRL(game Game, x int, y int) []string {
	returnSlice := []string{}
	returnSlice = append(returnSlice, game.board[x][y], game.board[x-1][y-1], game.board[x-2][y-2], game.board[x-3][y-3])
	return returnSlice
}


