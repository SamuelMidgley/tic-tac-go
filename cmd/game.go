package cmd

import "fmt"

func move(ai bool, oldMoves []string, numMoves int) []string {
	newMoves := oldMoves[:]

	player := getPlayer(numMoves)

	if ai && player == "x" {
		bestIndex := bestMove(newMoves)
		newMoves[bestIndex] = player
		return newMoves
	}

	playerMove := getInput(player, newMoves)
	newMoves[playerMove] = player

	return newMoves
}

func play(ai bool) {
	var result bool
	whoIsWinner := ""

	// Initialize moves
	numMoves := 0
	allMoves := make([]string, 9)

	for i := range allMoves {
		allMoves[i] = " "
	}

	intro()

	// Main game loop
	for numMoves <= 8 {
		allMoves = move(ai, allMoves, numMoves)

		numMoves++

		if numMoves >= 5 {
			// avoids unnecessary computations
			result, whoIsWinner = checkWin(allMoves)
		}

		buildBoard(allMoves)

		if result {
			break
		}
	}

	if result && whoIsWinner != "tie" {
		fmt.Printf("\nWell done %s, You Won!!", whoIsWinner)
	} else {
		fmt.Printf("\nUnlucky it's a draw")
	}
}
