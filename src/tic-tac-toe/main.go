package main

import (
	"fmt"
)

// returns a 3 by 3 (or 9 length) array as a memory address
func MakeBoard() *[9]int {
	var board [9]int

	for i := range board {
		// for every element in the range of the board
		// initialize it as 0
		board[i] = 0
	}

	return &board
}

// Player makes a play by putting 1 on a blank cell
func Play(board *[9]int, cell int) {
	board[cell] = 1
}

// CPU makes a play by putting 2 on a blank cell
func PlayCPU(board *[9]int) {
	var avails []int

	for i := range board {
		if board[i] == 0 {
			avails = append(avails, i)
		}
	}

	board[avails[len(avails)/2]] = 2
}

// verified if the cell on the board is not yet played
func Verify(board *[9]int, cell int) bool {
	if cell >= 0 && cell <= 9 {
		for i := range board {
			if i == cell && board[i] == 0 {
				return true
			}
		}
	}

	return false
}

// checks if somebody has won the game
func CheckBoard(board [9]int) int {
	// check for win
	wins := [8][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	for _, combo := range wins {
		if board[combo[0]] != 0 && board[combo[0]] == board[combo[1]] && board[combo[1]] == board[combo[2]] {
			return board[combo[0]]
		}
	}

	// if no win, check if it is a draw
	for i := range board {
		if board[i] == 0 {
			return -1
		}
	}

	return -2
}

// displays the board
func ShowBoard(board [9]int) {
	fmt.Printf("\n\nGame Board\n")
	fmt.Printf("\n%v..%v..%v\n%v..%v..%v\n%v..%v..%v\n\n", board[0], board[1], board[2], board[3], board[4], board[5], board[6], board[7], board[8])
}

// game loop
func GameLoop(board *[9]int) {
	var (
		name   string
		choice int
	)
	fmt.Print("What is your name: ")
	fmt.Scan(&name)

	fmt.Println("Welcome ", name, "!")

	for {
		ShowBoard(*board)

		fmt.Print("What is your move? Choice: ")
		fmt.Scan(&choice)

		if Verify(board, choice) {
			Play(board, choice)
		} else {
			fmt.Println("Invalid move. You're skipped!")
		}

		PlayCPU(board)

		check := CheckBoard(*board)
		if check == 1 {
			fmt.Println("Game end!", name, "won the game!")
			ShowBoard(*board)
			break
		} else if check == 2 {
			fmt.Println("Game end! CPU won the game!")
			ShowBoard(*board)
			break
		} else if check == -2 {
			fmt.Println("Game draw!")
			ShowBoard(*board)
			break
		}
	}
}

func main() {
	fmt.Println("Welcome to the tic-tac-toe game!")
	fmt.Println("You will play against a CPU.")

	var (
		board *[9]int = MakeBoard()
	)

	GameLoop(board)
}
