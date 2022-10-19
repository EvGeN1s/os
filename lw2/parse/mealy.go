package parse

import (
	"os/lw2/model"
)

func MealyToTable(a model.Mealy) [][]string {

	var result [][]string

	var states []string
	states = append(states, "")
	states = append(states, a.States...)
	result = append(result, states)

	for _, move := range a.Moves {
		moveWithStates := []string{move}
		for _, state := range a.States {
			stateWithSignal := a.StateMoveToStateSignal[state+move]
			moveWithStates = append(moveWithStates, stateWithSignal)
		}
		result = append(result, moveWithStates)
	}

	return result
}

func TableToMealy(records [][]string) model.Mealy {
	states := make([]string, 0, len(records[0])-1)

	for i, state := range records[0] {
		if i == 0 {
			continue
		}
		states = append(states, state)
	}

	moves := make([]string, 0, len(records)-1)
	stateMoveToState := make(map[string]string)

	for i, columns := range records {
		if i == 0 {
			continue
		}

		move := columns[0]
		moves = append(moves, move)

		for j, column := range columns {
			if j == 0 {
				continue
			}

			stateMoveToState[states[j-1]+move] = column
		}
	}

	return model.Mealy{
		Moves:                  moves,
		States:                 states,
		StateMoveToStateSignal: stateMoveToState,
	}
}
