package parse

import (
	"fmt"
	"os/lw3/grammatictoautomat/model"
)

func AutomateToTable(automate model.Automate) [][]string {
	newNames := make(map[string]string)
	res := make([][]string, 0)

	var states []string
	states = append(states, "")
	for i, state := range automate.States {
		newName := fmt.Sprintf("S%d", i)
		newNames[state] = newName

		fmt.Printf("%s = %s\n", newName, state)

		states = append(states, newName)
	}

	res = append(res, states)

	for _, move := range automate.Moves {
		row := []string{move}
		for _, state := range automate.States {
			stateMove := state + move
			if _, found := automate.StateMoveToState[stateMove]; found {
				resState := automate.StateMoveToState[stateMove]
				row = append(row, newNames[resState])
			} else {
				row = append(row, "")
			}
		}

		res = append(res, row)
	}

	return res
}
