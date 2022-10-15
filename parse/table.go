package parse

import (
	"fmt"
	"os/model"
	"os/util"
	"strings"
)

const (
	MuraOutStateChar  = "s"
	MilliOutStateChar = "a"

	MoveChar   = "z"
	SignalChar = "y"
)

func ConvertFromTableToMilli(records [][]string) model.Milli {
	result := make(model.Milli, len(records)-1, len(records)-1)

	states := make([]int, 0, len(records[0])-1)

	for i, state := range records[0] {
		if i == 0 {
			continue
		}
		states = append(states, util.StringToInt(state))
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
				result[move] = make(map[int]model.MilliState)
			}

			result[move][states[j-1]] = state
		}
	}

	return result
}

func ConvertFromTableToMura(records [][]string) model.Mura {
	result := make(model.Mura, len(records)-1, len(records)-1)

	signals := make([]int, 0, len(records[0])-1)
	for i, sign := range records[0] {
		if i == 0 {
			continue
		}
		signals = append(signals, util.StringToInt(sign))
	}

	states := make([]int, 0, len(records[0])-1)
	for i, state := range records[1] {
		if i == 0 {
			continue
		}
		states = append(states, util.StringToInt(state))
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
			state := util.StringToInt(column)

			muraState := model.MuraState{
				State:  states[j-1],
				Signal: signals[j-1],
			}

			result[move][muraState] = state
		}
	}

	return result
}

func ConvertFromMuraToTable(mura model.Mura) [][]string {
	var result [][]string
	signals := make([]string, 0)
	signals = append(signals, "")
	states := make([]string, 0)
	states = append(states, "")
	intState := make([]model.MuraState, 0)
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

func ConvertFromMilliToTable(milli model.Milli) [][]string {
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

func convertStringToState(s string) model.MilliState {
	res := strings.Split(s, "/")

	return model.MilliState{
		State:  util.StringToInt(res[0]),
		Signal: util.StringToInt(res[1]),
	}
}

func covertStateToString(state model.MilliState) string {
	return fmt.Sprintf("%s%d/%s%d", MilliOutStateChar, state.State, SignalChar, state.Signal)
}
