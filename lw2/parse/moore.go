package parse

import (
	"os/lw2/model"
)

const MooreMoveRowIndex = 2

func TableToMoore(records [][]string) model.Moore {
	states := make([]string, 0, len(records[0])-1)
	stateToSignal := make(map[string]string)

	for i := 1; i < len(records[0]); i++ {
		signal := records[0][i]
		state := records[1][i]

		stateToSignal[state] = signal

		states = append(states, state)
	}

	moves, stateMoveToState := getStateMoveToState(records, states, MooreMoveRowIndex)

	return model.Moore{
		Moves:            moves,
		States:           states,
		StateToSignal:    stateToSignal,
		StateMoveToState: stateMoveToState,
	}
}

func MooreToTable(a model.Moore) [][]string {
	var result [][]string
	signals := make([]string, 0)
	signals = append(signals, "")
	states := make([]string, 0)
	states = append(states, "")

	for _, state := range a.States {
		signal := a.StateToSignal[state]
		signals = append(signals, signal)
		states = append(states, state)
	}

	result = append(result, signals)
	result = append(result, states)

	for _, move := range a.Moves {
		var statesWithMoves []string
		statesWithMoves = append(statesWithMoves, move)
		for _, state := range a.States {
			statesWithMoves = append(statesWithMoves, a.StateMoveToState[state+move])
		}

		result = append(result, statesWithMoves)
	}

	return result
}

func getStateMoveToState(records [][]string, states []string, moveIndex int) ([]string, map[string]string) {
	moves := make([]string, 0, len(records)-1)
	stateMoveToState := make(map[string]string)

	for i, columns := range records {
		if i < moveIndex {
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

	return moves, stateMoveToState
}
