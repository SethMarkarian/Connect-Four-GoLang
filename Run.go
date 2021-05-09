package main

import "fmt"

func main() {
	g := createGame()
	printBoard(g)
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
			} else if g.piecesInColumn[place-1] >= 6 {
				fmt.Println("Column is full, try another column")
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
			} else if g.piecesInColumn[place-1] >= 6 {
				fmt.Println("Column is full, try another column")
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
		} else if boardFilled(g) {
			g.done = true
		}
	}
	return
}
