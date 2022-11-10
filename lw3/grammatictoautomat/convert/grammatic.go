package convert

import (
	"container/list"
	"fmt"
	"os/lw3/grammatictoautomat/model"
)

func GrammaticToAutomate(grammatic model.Grammatic) model.Automate {
	queue := list.New()

	var resStates []string
	resStateMoveToState := make(map[string]string)
	checkedStates := make(map[string]bool)
	checkedPrevSates := make(map[string]bool)

	for _, state := range grammatic.States {
		if _, found := checkedPrevSates[state]; found {
			continue
		}
		queue.PushBack(state)

		for queue.Len() > 0 {
			curr := queue.Front()
			currState := fmt.Sprintf("%v", curr.Value)

			var newStates []string
			stateMoveToState := make(map[string]string)
			newStates, stateMoveToState, checkedPrevSates = getStates(currState, grammatic.StateMoveToStates, grammatic.Moves, checkedPrevSates)
			for _, newState := range newStates {
				if _, found := checkedStates[newState]; found {
					continue
				}

				checkedStates[newState] = true
				if newState == currState {
					continue
				}

				queue.PushBack(newState)
			}

			resStates = append(resStates, currState)
			for stateMove, resState := range stateMoveToState {
				resStateMoveToState[stateMove] = resState
			}
			checkedStates[currState] = true

			queue.Remove(curr)
		}
	}

	return model.Automate{
		Moves:            grammatic.Moves,
		States:           resStates,
		StateMoveToState: resStateMoveToState,
	}
}

func getStates(currState string, stateMoveToStates map[string][]string, moves []string, checkedPrevStates map[string]bool) ([]string, map[string]string, map[string]bool) {
	resStateMoveToStates := make(map[string]string)
	var newStates []string
	for _, move := range moves {
		var newState string
		for _, stateRune := range currState {
			state := fmt.Sprintf("%c", stateRune)
			if _, found := stateMoveToStates[state+move]; !found {
				continue
			}

			for _, resState := range stateMoveToStates[state+move] {
				checkedPrevStates[resState] = true
				newState += resState
			}
		}

		if newState == "" {
			continue
		}
		resStateMoveToStates[currState+move] = newState
		newStates = append(newStates, newState)
	}

	return newStates, resStateMoveToStates, checkedPrevStates
}
