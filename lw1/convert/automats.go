package convert

import (
	model2 "os/lw1/model"
	util2 "os/lw1/util"
)

func MilliToMure(milli model2.Milli) model2.Mura {
	result := make(model2.Mura, len(milli), len(milli))

	newStates, muraToMilliState, milliToMuraState := findCommonMuraState(milli)

	for _, state := range newStates {
		milliState := muraToMilliState[state.State]

		for move, states := range milli {
			if result[move] == nil {
				result[move] = make(map[model2.MuraState]int)
			}

			moveState := states[milliState.State]
			muraState := milliToMuraState[moveState]

			result[move][state] = muraState
		}
	}

	return result
}

func MuraToMilli(mura model2.Mura) model2.Milli {
	result := make(model2.Milli, len(mura), len(mura))

	statesToSign := make(map[int]int)
	for _, states := range mura {
		for state := range states {
			statesToSign[state.State] = state.Signal
		}
	}

	for move, states := range mura {
		for inState, state := range states {
			milliState := model2.MilliState{
				State:  state,
				Signal: statesToSign[state],
			}
			if result[move] == nil {
				result[move] = make(map[int]model2.MilliState)
			}

			result[move][inState.State] = milliState
		}
	}

	return result
}

func findCommonMuraState(milli model2.Milli) ([]model2.MuraState, map[int]model2.MilliState, map[model2.MilliState]int) {
	var result []model2.MuraState
	var uniqStates []model2.MilliState
	muraToMilliMap := make(map[int]model2.MilliState)
	milliToMuraMap := make(map[model2.MilliState]int)
	statesCount := 0
	for _, states := range milli {
		for _, state := range states {
			newState := model2.MuraState{
				State:  statesCount,
				Signal: state.Signal,
			}

			if !util2.ContainsState(uniqStates, state) {
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
