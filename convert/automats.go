package convert

import (
	"os/model"
	"os/util"
)

func ConvertMilliToMure(milli model.Milli) model.Mura {
	result := model.Mura{}

	newStates := findCommonMuraState(milli)
	for _, state := range newStates {
		for move, states := range milli {
			result[move][state] = states[state.State].State
		}
	}

	return result
}

func ConvertMuraToMilli(mura model.Mura) model.Milli {
	result := model.Milli{}

	statesToSign := make(map[int]int)
	for _, states := range mura {
		for state := range states {
			statesToSign[state.State] = state.Signal
		}
	}

	for move, states := range mura {
		for inState, state := range states {
			milliState := model.MilliState{
				State:  state,
				Signal: statesToSign[state],
			}

			result[move][inState.State] = milliState
		}
	}

	return result
}

func findCommonMuraState(milli model.Milli) []model.MuraState {
	var result []model.MuraState
	for _, states := range milli {
		for _, state := range states {
			newState := model.MuraState{
				State:  state.State,
				Signal: state.Signal,
			}

			if !util.ContainsMuraState(result, newState) {
				result = append(result, newState)
			}
		}
	}

	return result
}
