package main

import (
	"fmt"
	"strconv"
)

func intro() {
	// Introduction to the game and how to play
	fmt.Println("\nWelcome to Tic Tac Go!!")
	fmt.Println("When you are asked for a move provide the corresponding index you want")
	fmt.Println("See the board below for guidance")

	// Create reference index board
	guideMoves := make([]string, 9)
	index := 0
	for i := range guideMoves {
		guideMoves[i] = strconv.Itoa(index)
		index++
	}

	buildBoard(guideMoves)
}

func getMove(oldMoves []string, numMoves int) (newMoves []string) {
	newMoves = oldMoves[:]
	var index int
	var player string

	// Determine which players' turn it is
	if numMoves%2 == 0 {
		player = "x"
	} else {
		player = "o"
	}

	// Handles player input, includes error handling for out of bounds index
	// and for wrong type
	var s string
	for {
		fmt.Printf("\n%s, What is your move? ", player)
		_, err := fmt.Scan(&s)
		index, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println("Please enter an integer from 0-8")
		} else {
			if index >= 0 && index <= 8 {
				if newMoves[index] == " " {
					newMoves[index] = player
					break
				} else {
					fmt.Println("Sorry, that move is taken")
				}
			} else {
				fmt.Println("Please enter an integer from 0-8")
			}
		}
	}
	return
}

func buildBoard(allMoves []string) {
	fmt.Printf("\n %s | %s | %s \n", allMoves[0], allMoves[1], allMoves[2])
	fmt.Println("---+---+---")
	fmt.Printf(" %s | %s | %s \n", allMoves[3], allMoves[4], allMoves[5])
	fmt.Println("---+---+---")
	fmt.Printf(" %s | %s | %s \n", allMoves[6], allMoves[7], allMoves[8])
}

func checkWin(allMoves []string) (result bool, player string) {
	// Determines whether the game is complete through brute force
	var winningIndexes = [8][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	var value int
	for _, indexes := range winningIndexes {
		total := 0
		for _, index := range indexes {
			switch allMoves[index] {
			case "x":
				value = 1
			case "o":
				value = -1
			default:
				value = 0
			}
			total += value
		}
		if total == 3 {
			result = true
			player = "x"
			break
		}
		if total == -3 {
			result = true
			player = "o"
			break
		} else {
			result = false
		}
	}
	return
}

func main() {
	var result bool
	whoIsWinner := ""

	// Initialise moves
	numMoves := 0
	allMoves := make([]string, 9)

	for i := range allMoves {
		allMoves[i] = " "
	}

	intro()

	// Main game loop
	for numMoves < 10 {
		allMoves = getMove(allMoves, numMoves)
		numMoves++

		if numMoves >= 5 {
			// avoids unneccesary computations
			result, whoIsWinner = checkWin(allMoves)
		}

		buildBoard(allMoves)

		if result {
			break
		}
	}

	if result {
		fmt.Printf("\nWell done %s, You Won!!", whoIsWinner)
	} else {
		fmt.Printf("\nUnlucky it's a draw")
	}
}
