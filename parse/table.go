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
	result := model.Milli{}

	states := make([]int, 0, len(records[0])-1)
	for i := 1; i < len(records); i++ {
		states = append(states, util.StringToInt(records[i][0]))
	}

	/*for i, state := range records[0] {
		if i == 0 {
			continue
		}
		states = append(states, util.StringToInt(state))
	}*/

	for i, rows := range records {
		if i == 0 {
			continue
		}
		move := i - 1
		for j, column := range rows {
			if j == 0 {
				continue
			}
			state := convertState(column)

			result[move][states[j-1]] = state
		}
	}

	return result
}

func ConvertFromTableToMura(records [][]string) model.Mura {
	result := model.Mura{}

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

func ConvertMilliToOutStr(milli model.Milli) string {
	result := ";"
	for _, inStates := range milli {
		for inState := range inStates {
			result += fmt.Sprintf("%s%d;", MilliOutStateChar, inState)
		}
	}

	for move, inStates := range milli {
		result += fmt.Sprintf("%s%d;", MoveChar, move)
		for _, state := range inStates {
			result += fmt.Sprintf("%s%d/%s%d;", MilliOutStateChar, state.State, SignalChar, state.Signal)
		}
	}

	return result
}

func ConvertMuraToOutStr(mura model.Mura) string {
	result := ";"
	for _, inStates := range mura {
		for inState := range inStates {
			result += fmt.Sprintf("%s%d;", SignalChar, inState.Signal)
		}
	}
	result += ";"
	for _, inStates := range mura {
		for inState := range inStates {
			result += fmt.Sprintf("%s%d;", MuraOutStateChar, inState.State)
		}
	}

	for move, inStates := range mura {
		result += fmt.Sprintf("%s%d;", MoveChar, move)
		for _, state := range inStates {
			result += fmt.Sprintf("%s%d;", MuraOutStateChar, state)
		}
	}

	return result
}

func convertState(s string) model.MilliState {
	res := strings.Split(s, "/")

	return model.MilliState{
		State:  util.StringToInt(res[0]),
		Signal: util.StringToInt(res[1]),
	}
}
