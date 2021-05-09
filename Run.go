package main

import "fmt"

func main() {
	g := createGame() //Creates new instance of a blank Game
	printBoard(g)
	numTurns := 0
	for !g.done { // Loops until a win or draw
		place := 0
		if numTurns%2 == 0 { //Player 1 turn
			fmt.Println("Player 1 turn")
			fmt.Print("Column: ")
			fmt.Scan(&place)
			if place > 7 { //User inputs a number outside of the board
				fmt.Println("Enter a column from 1-7")
				continue
			} else if g.piecesInColumn[place-1] >= 6 { //User inputs a column that is filled already
				fmt.Println("Column is full, try another column")
				continue
			}
			g = addPiece(place-1, "X", g) // Adds a 'X' piece at inputted column
		} else { // Player 2 turn
			fmt.Println("Player 2 turn")
			fmt.Print("Column: ")
			fmt.Scan(&place)
			if place > 7 { //User inputs a number outside of the board
				fmt.Println("Enter a column from 1-7")
				continue
			} else if g.piecesInColumn[place-1] >= 6 { //User inputs a column that is filled already
				fmt.Println("Column is full, try another column")
				continue
			}
			g = addPiece(place-1, "O", g) // Adds a 'O' piece at inputted column
		}
		printBoard(g)
		fmt.Println()
		numTurns = numTurns + 1
		player := gameWin(g) //Checks if updated game has a win
		if player == 1 {
			fmt.Println("Player 1 Wins!")
			break
		} else if player == 2 {
			fmt.Println("Player 2 Wins!")
			break
		} else if boardFilled(g) {
			g.done = true
		}
	}
	return
}
