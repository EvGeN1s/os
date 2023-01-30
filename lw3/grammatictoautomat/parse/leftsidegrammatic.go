package parse

import (
	"fmt"
	"os/lw3/grammatictoautomat/model"
	"regexp"
	"strings"
)

func RowsToRGrammatic(rows []string) model.Grammatic {
	var states []string
	checkedStates := make(map[string]bool)

	var moves []string
	checkedMoves := make(map[string]bool)

	stateMoveToStates := make(map[string][]string)

	for _, row := range rows {
		r := regexp.MustCompile("\\s+")
		row = r.ReplaceAllString(row, "")

		stateWithStateMoves := strings.Split(row, "->")
		state := stateWithStateMoves[0]
		if _, found := checkedStates[state]; !found {
			states = append(states, state)
			checkedStates[state] = true
		}

		stateMoves := strings.Split(stateWithStateMoves[1], "|")

		for _, stateMove := range stateMoves {
			move := fmt.Sprintf("%c", stateMove[0])
			if _, found := checkedMoves[move]; !found {
				moves = append(moves, move)
				checkedMoves[move] = true
			}

			var resState string
			if len(stateMove) == 1 {
				resState = "f"
			} else {
				resState = fmt.Sprintf("%c", stateMove[1])
			}

			i := state + move
			if stateMoveToStates[i] == nil {
				stateMoveToStates[i] = make([]string, 0)
			}

			currResStates := stateMoveToStates[i]
			currResStates = append(currResStates, resState)
			stateMoveToStates[i] = currResStates

			if _, found := checkedStates[resState]; !found {
				states = append(states, resState)
				checkedStates[resState] = true
			}
		}
	}

	return model.Grammatic{
		States:            states,
		Moves:             moves,
		StateMoveToStates: stateMoveToStates,
	}
}
