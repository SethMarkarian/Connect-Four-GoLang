package main

import (
	"fmt"
)

type Game struct {
	board [][]string
	done  bool
}

func createGame() Game {
	newGame := Game{}
	newGame.board = [][]string{
		{"_", "_", "_", "_", "_", "_", "_"}, {"_", "_", "_", "_", "_", "_", "_"}, {"_", "_", "_", "_", "_", "_", "_"}, {"_", "_", "_", "_", "_", "_", "_"}, {"_", "_", "_", "_", "_", "_", "_"}, {"_", "_", "_", "_", "_", "_", "_"},
	}
	newGame.done = false
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
	for i := 0; i < 6; i++ {
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

func main() {
	g := createGame()
	numTurns := 0
	for !g.done {
		place := 0
		if numTurns%2 == 0 {
			fmt.Println("Player 1 turn")
			fmt.Print("Column: ")
			fmt.Scan(&place)
			if place > 7 {
				fmt.Println("Enter a column from 1-7")
				continue
			}
			g = addPiece(place-1, "X", g)
		} else {
			fmt.Println("Player 2 turn")
			fmt.Print("Column: ")
			fmt.Scan(&place)
			if place > 7 {
				fmt.Println("Enter a column from 1-7")
				continue
			}
			g = addPiece(place-1, "O", g)
		}
		printBoard(g)
		fmt.Println()
		numTurns = numTurns + 1
		player := gameWin(g)
		if player == 1 {
			fmt.Println("Player 1 Wins!")
			break
		} else if player == 2 {
			fmt.Println("Player 2 Wins!")
			break
		}
	}
	return
}
