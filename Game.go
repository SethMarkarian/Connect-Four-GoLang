package main

import (
	"fmt"
)

// Creates Game structure
type Game struct {
	board          [][]string // Slice that holds '_', 'X', or 'O'
	piecesInColumn []int      // Slice that counts amount of pieces in given column
	done           bool       // Gives status of game, false if no, true if yes
}

// Initializes all variables held in Game
// Fills board with '_'
// Returns a Game
func createGame() Game {
	newGame := Game{}
	newGame.board = [][]string{
		{"_", "_", "_", "_", "_", "_", "_"}, {"_", "_", "_", "_", "_", "_", "_"}, {"_", "_", "_", "_", "_", "_", "_"}, {"_", "_", "_", "_", "_", "_", "_"}, {"_", "_", "_", "_", "_", "_", "_"}, {"_", "_", "_", "_", "_", "_", "_"},
	}
	newGame.done = false
	newGame.piecesInColumn = []int{0, 0, 0, 0, 0, 0, 0}
	return newGame
}

// Prints board in a user-friendly format
func printBoard(game Game) {
	for i := 0; i < 6; i++ {
		fmt.Println(game.board[i])
	}
	fmt.Println(" 1 2 3 4 5 6 7")
}

// Returns a boolean value if there is a draw (cannot add anymore pieces to the board)
// Returns true if it's a draw, false if not
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

// Adds a piece to the board given the piece and column (Simulates piece being "dropped")
// Returns an updated Game with piece added
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

//Checks if there is 4 same pieces in a row, column, or diagonal
//Returns 1 (Player 1 wins), 2 (Player 2 wins), -1 (No one wins)
func gameWin(game Game) int {
	// Horizontal check
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

	// vertical check
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

	// Diagonal L-R check
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

	// Diagonal R-L check
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

// Helper function to compare a slice and a string to see if all elements in the slice are equal to the string
// Returns true if yes, false if no
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

// Helper function that adds a given column into a slice
// Returns a slice that represents the column at number "col"
func getVerticalLine(game Game, col int) []string {
	returnSlice := []string{}
	for i := 0; i < 6; i++ {
		returnSlice = append(returnSlice, game.board[i][col])
	}
	return returnSlice
}

// Helper function that adds a given LR diagonal into a slice given an x and y start point
// Returns a slice that represents a LR diagonal that starts at the given start point
func getDiagLR(game Game, x int, y int) []string {
	returnSlice := []string{}
	returnSlice = append(returnSlice, game.board[x][y], game.board[x-1][y+1], game.board[x-2][y+2], game.board[x-3][y+3])
	return returnSlice
}

// Helper function that adds a given RL diagonal into a slice given an x and y start point
// Returns a slice that represents a RL diagonal that starts at the given start point
func getDiagRL(game Game, x int, y int) []string {
	returnSlice := []string{}
	returnSlice = append(returnSlice, game.board[x][y], game.board[x-1][y-1], game.board[x-2][y-2], game.board[x-3][y-3])
	return returnSlice
}
