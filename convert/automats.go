package convert

import (
	"os/model"
	"os/util"
)

func MilliToMure(milli model.Milli) model.Mura {
	result := make(model.Mura, len(milli), len(milli))

	newStates, muraToMilliState, milliToMuraState := findCommonMuraState(milli)

	for _, state := range newStates {
		milliState := muraToMilliState[state.State]

		for move, states := range milli {
			if result[move] == nil {
				result[move] = make(map[model.MuraState]int)
			}

			moveState := states[milliState.State]
			muraState := milliToMuraState[moveState]

			result[move][state] = muraState
		}
	}

	return result
}

func MuraToMilli(mura model.Mura) model.Milli {
	result := make(model.Milli, len(mura), len(mura))

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
			if result[move] == nil {
				result[move] = make(map[int]model.MilliState)
			}

			result[move][inState.State] = milliState
		}
	}

	return result
}

func findCommonMuraState(milli model.Milli) ([]model.MuraState, map[int]model.MilliState, map[model.MilliState]int) {
	var result []model.MuraState
	var uniqStates []model.MilliState
	muraToMilliMap := make(map[int]model.MilliState)
	milliToMuraMap := make(map[model.MilliState]int)
	statesCount := 0
	for _, states := range milli {
		for _, state := range states {
			newState := model.MuraState{
				State:  statesCount,
				Signal: state.Signal,
			}

			if !util.ContainsState(uniqStates, state) {
				uniqStates = append(uniqStates, state)

				result = append(result, newState)
				muraToMilliMap[statesCount] = state
				milliToMuraMap[state] = statesCount
				statesCount++
			}
		}
	}

	return result, muraToMilliMap, milliToMuraMap
}
