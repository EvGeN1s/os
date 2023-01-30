package parse

import (
	"os/lw3/automatenormaliztion/model"
	"strings"
)

const tableStart = 2

func TableToMultiAuto(records [][]string) model.MultiStateAuto {
	states := make([]string, 0, len(records[0])-1)
	finalStates := make(map[string]bool)

	for i := 1; i < len(records[0]); i++ {
		signal := records[0][i]
		state := records[1][i]

		if signal == "F" {
			finalStates[state] = true
		}

		states = append(states, state)
	}

	moves, stateMoveToStates := getStateMoveToStates(records, states)

	return model.MultiStateAuto{
		States:           states,
		Moves:            moves,
		FinalStates:      finalStates,
		StateMoveToState: stateMoveToStates,
	}
}

func getStateMoveToStates(records [][]string, states []string) ([]string, map[string][]string) {
	moves := make([]string, 0, len(records)-1)
	stateMoveToState := make(map[string][]string)

	for i, columns := range records {
		if i < tableStart {
			continue
		}

		move := columns[0]
		moves = append(moves, move)

		for j, column := range columns {
			if j == 0 {
				continue
			}
			if column == "" {
				continue
			}

			resStates := strings.Split(column, ",")
			stateMoveToState[states[j-1]+move] = resStates
		}
	}

	return moves, stateMoveToState
}
