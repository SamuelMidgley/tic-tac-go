package cmd

func bestMove(allMoves []string) (move int) {
	prev_value := " "
	bestScore := -1000
	for index, value := range allMoves {
		if value == " " {
			if prev_value == "x" {
				allMoves[index] = "o"
			} else {
				allMoves[index] = "x"
			}
			prev_value = value

			score := minimax(allMoves, 0, false)
			allMoves[index] = " "

			if bestScore < score {
				bestScore = score
				move = index
			}
		}
	}
	return
}

func minimax(allMoves []string, depth int, isMaximizing bool) (bestScore int) {
	result, player := checkWin(allMoves)
	if result {
		convertPlayer := map[string]int{
			"x": 1, "o": -1, "tie": 0,
		}
		bestScore = convertPlayer[player]
		return
	}
	if isMaximizing {
		bestScore = -1000
		for index, value := range allMoves {
			if value == " " {
				allMoves[index] = "x"
				score := minimax(allMoves, depth+1, false)
				allMoves[index] = " "
				if bestScore < score {
					bestScore = score
				}
			}
		}
		return
	} else {
		bestScore = 1000
		for index, value := range allMoves {
			if value == " " {
				allMoves[index] = "o"
				score := minimax(allMoves, depth+1, true)
				allMoves[index] = " "
				if bestScore > score {
					bestScore = score
				}
			}
		}
		return
	}
}
