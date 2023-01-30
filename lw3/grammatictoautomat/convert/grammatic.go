package convert

import (
	"container/list"
	"fmt"
	"os/lw3/grammatictoautomat/model"
)

func GrammaticToAutomate(grammatic model.Grammatic, numsCount int) model.Automate {
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
			newStates, stateMoveToState, checkedPrevSates = getStates(currState, grammatic.StateMoveToStates, grammatic.Moves, checkedPrevSates, numsCount)
			fmt.Println(newStates)

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

func getStates(currState string, stateMoveToStates map[string][]string, moves []string, checkedPrevStates map[string]bool, numsCount int) ([]string, map[string]string, map[string]bool) {
	resStateMoveToStates := make(map[string]string)
	var newStates []string
	for _, move := range moves {
		var newState string
		checkedStates := make(map[string]bool)
		var state string
		for i, stateRune := range currState {
			if i%numsCount == 0 && numsCount > 1 {
				state = fmt.Sprintf("%c", stateRune)
				continue
			} else {
				state = fmt.Sprintf("%s%c", state, stateRune)
			}
			if _, found := stateMoveToStates[state+move]; !found {
				continue
			}

			for _, resState := range stateMoveToStates[state+move] {
				var outputState string
				var state1 string
				for j, stateRune1 := range resState {
					if j%numsCount == 0 && numsCount > 1 {
						state1 = fmt.Sprintf("%c", stateRune1)
						continue
					} else {
						state1 = fmt.Sprintf("%s%c", state1, stateRune1)
					}

					if _, found := checkedStates[state1]; found {
						continue
					}

					outputState += state1
					checkedStates[state1] = true
					checkedPrevStates[state1] = true
				}

				newState += outputState
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
