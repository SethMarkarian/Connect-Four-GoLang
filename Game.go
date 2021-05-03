package main

import "fmt"

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
		// winX := []string{"X", "X", "X", "X"}
		for i := 0; i <= len(line)-4; i++ {
			if compare(line[i:i+4], "X") {
				game.done = true
				return 1
			} else if compare(line[i:i+4], "O") {
				game.done = true
				return 2
			}
		}
	}
	return -1
}

func compare(slice []string, stone string) bool {
	if slice == nil {
		return false
	}
	for _, char := range slice {
		if char != stone {
			return false
		}
	}
	return true
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
			g = addPiece(place-1, "X", g)
		} else {
			fmt.Println("Player 2 turn")
			fmt.Print("Column: ")
			fmt.Scan(&place)
			g = addPiece(place-1, "O", g)
		}
		printBoard(g)
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
}
