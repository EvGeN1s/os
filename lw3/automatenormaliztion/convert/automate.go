package convert

import (
	model2 "os/lw3/automatenormaliztion/model"
	"os/lw3/grammatictoautomat/model"
)

func RemoveEmptyMoves(automate model2.MultiStateAuto) model.Grammatic {
	stateMoveToStates := make(map[string][]string)
	states := make([]string, 0)

	for _, state := range automate.States {
		checkedStates := make(map[string]bool)
		resStates := []string{state}

		resStates = getNewState(state, automate.StateMoveToState, checkedStates, resStates, automate.Moves[len(automate.Moves)-1])

		var resState string
		for _, s := range resStates {
			resState += s
		}

		for _, s := range resStates {
			for i, move := range automate.Moves {
				// skip E sign
				if i == len(automate.Moves)-1 {
					continue
				}

				stateMove := s + move
				resStateMove := state + move

				if foundStates, found := automate.StateMoveToState[stateMove]; found && len(foundStates) != 0 {
					if _, ok := stateMoveToStates[resStateMove]; !ok {
						stateMoveToStates[resStateMove] = make([]string, 0)
					}

					//fmt.Println( "State move:", resStateMove, "Curr States: ", stateMoveToStates[resStateMove], "Add states: ", foundStates, len(foundStates))
					stateMoveToStates[resStateMove] = merge(stateMoveToStates[resStateMove], foundStates)
					//fmt.Println("Result: ", stateMoveToStates[resStateMove])
				}
			}
		}

		states = append(states, resState)
	}

	return model.Grammatic{
		States:            automate.States,
		Moves:             automate.Moves[:len(automate.Moves)-1],
		StateMoveToStates: stateMoveToStates,
	}
}

func getNewState(currState string, stateMoveToState map[string][]string, checkedStates map[string]bool, states []string, move string) []string {
	if foundStates, found := stateMoveToState[currState+move]; found {
		for _, foundState := range foundStates {
			if _, foundRes := checkedStates[foundState]; foundRes {
				return states
			}
			checkedStates[foundState] = true

			states = append(states, foundState)
			states = getNewState(foundState, stateMoveToState, checkedStates, states, move)
		}
	}
	return states
}

func merge(res, exp []string) []string {
	resValues := make(map[string]bool)
	for _, val := range res {
		resValues[val] = true
	}

	for _, val := range exp {
		if found := resValues[val]; !found {
			res = append(res, val)
		}
	}

	return res
}
