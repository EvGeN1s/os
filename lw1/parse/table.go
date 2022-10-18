package parse

import (
	"fmt"
	model2 "os/lw1/model"
	util2 "os/lw1/util"
	"strings"
)

const (
	MuraOutStateChar  = "s"
	MilliOutStateChar = "a"

	MoveChar   = "z"
	SignalChar = "y"
)

func ConvertFromTableToMilli(records [][]string) model2.Milli {
	result := make(model2.Milli, len(records)-1, len(records)-1)

	states := make([]int, 0, len(records[0])-1)

	for i, state := range records[0] {
		if i == 0 {
			continue
		}
		states = append(states, util2.StringToInt(state))
	}

	for i, rows := range records {
		if i == 0 {
			continue
		}
		move := i - 1
		for j, column := range rows {
			if j == 0 {
				continue
			}
			state := convertStringToState(column)
			if result[move] == nil {
				result[move] = make(map[int]model2.MilliState)
			}

			result[move][states[j-1]] = state
		}
	}

	return result
}

func ConvertFromTableToMura(records [][]string) model2.Mura {
	result := make(model2.Mura, len(records)-2, len(records)-2)

	signals := make([]int, 0, len(records[0])-1)
	for i, sign := range records[0] {
		if i == 0 {
			continue
		}
		signals = append(signals, util2.StringToInt(sign))
	}

	states := make([]int, 0, len(records[0])-1)
	for i, state := range records[1] {
		if i == 0 {
			continue
		}
		states = append(states, util2.StringToInt(state))
	}

	for i, rows := range records {
		if i < 2 {
			continue
		}
		move := i - 2
		for j, column := range rows {
			if j == 0 {
				continue
			}
			state := util2.StringToInt(column)

			muraState := model2.MuraState{
				State:  states[j-1],
				Signal: signals[j-1],
			}

			if result[move] == nil {
				result[move] = make(map[model2.MuraState]int)
			}

			result[move][muraState] = state
		}
	}

	return result
}

func ConvertFromMuraToTable(mura model2.Mura) [][]string {
	var result [][]string
	signals := make([]string, 0)
	signals = append(signals, "")
	states := make([]string, 0)
	states = append(states, "")
	intState := make([]model2.MuraState, 0)
	for _, muraSates := range mura {
		for muraState := range muraSates {
			signals = append(signals, fmt.Sprintf("%s%d", SignalChar, muraState.Signal))
			states = append(states, fmt.Sprintf("%s%d", MuraOutStateChar, muraState.State))
			intState = append(intState, muraState)
		}
		break
	}

	result = append(result, signals)
	result = append(result, states)

	for move, muraStates := range mura {
		statesWithMoves := make([]string, 0, len(muraStates)+1)
		statesWithMoves = append(statesWithMoves, fmt.Sprintf("%s%d", MoveChar, move))
		for _, state := range intState {
			statesWithMoves = append(statesWithMoves, fmt.Sprintf("%s%d", MuraOutStateChar, muraStates[state]))
		}

		result = append(result, statesWithMoves)
	}

	return result
}

func ConvertFromMilliToTable(milli model2.Milli) [][]string {
	var result [][]string
	states := make([]string, 0)
	states = append(states, "")
	intState := make([]int, 0)

	for _, milliStates := range milli {
		for state := range milliStates {
			states = append(states, fmt.Sprintf("%s%d", MilliOutStateChar, state))
			intState = append(intState, state)
		}
		break
	}

	result = append(result, states)

	for move, milliStates := range milli {
		statesWithMoves := make([]string, 0, len(milliStates)+1)
		statesWithMoves = append(statesWithMoves, fmt.Sprintf("%s%d", MoveChar, move))
		for _, state := range intState {
			statesWithMoves = append(statesWithMoves, covertStateToString(milliStates[state]))
		}

		result = append(result, statesWithMoves)
	}

	return result
}

func convertStringToState(s string) model2.MilliState {
	res := strings.Split(s, "/")

	return model2.MilliState{
		State:  util2.StringToInt(res[0]),
		Signal: util2.StringToInt(res[1]),
	}
}

func covertStateToString(state model2.MilliState) string {
	return fmt.Sprintf("%s%d/%s%d", MilliOutStateChar, state.State, SignalChar, state.Signal)
}
