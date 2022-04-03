package main

import "fmt"

var board [3][3]int

func main() {

	for i := 0; i < 9; i++ {
	prompt:
		ret := askForMove(i%2 + 1)
		if ret < 0 || ret > 8 || board[ret/3][ret%3] != 0 {
			goto prompt
		}
		set(ret, i%2+1)
		if checkForWin() != 0 {
			break
		}
	}
	printBoard()
	if checkForWin() == 0 {
		fmt.Printf("Das Spiel endet unentschieden.\n")
	} else {
		fmt.Printf("Spieler %d hat gewonnen.\n", checkForWin())
	}
}

func set(move, player int) {
	board[move/3][move%3] = player
}

func askForMove(player int) (move int) {

	printBoard()
	fmt.Printf("Spieler %d, auf welches Feld willst du setzen?\n", player)
	fmt.Scanf("%d", &move)
	return move
}

func checkForWin() (winner int) {

	for player := 1; player < 3; player++ {

		for i := 0; i < 3; i++ {
			if board[i][0] == player && board[i][1] == player && board[i][2] == player {
				return player
			}
			if board[0][i] == player && board[1][i] == player && board[2][i] == player {
				return player
			}
		}
		if board[0][0] == player && board[1][1] == player && board[2][2] == player {
			return player
		}
		if board[0][2] == player && board[1][1] == player && board[2][0] == player {
			return player
		}
	}

	return 0
}

func printBoard() {

	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {

			if j < 2 {

				if board[i][j] == 0 {
					fmt.Printf("%d | ", i*3+j)
				} else if board[i][j] == 1 {
					fmt.Printf("%c | ", 'x')
				} else if board[i][j] == 2 {
					fmt.Printf("%c | ", 'o')
				}

			} else {

				if board[i][j] == 0 {
					fmt.Printf("%d\n", i*3+j)
				} else if board[i][j] == 1 {
					fmt.Printf("%c\n", 'x')
				} else if board[i][j] == 2 {
					fmt.Printf("%c\n", 'o')
				}
			}
		}

		if i < 2 {
			fmt.Printf("--+---+--\n")
		}
	}
}
