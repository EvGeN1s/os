package convert

import "os/lw2/model"

func MooreToEqualityGroup(a model.Moore) model.EqualityGroup {
	removeUnreachableMooreStates(&a)

	return model.EqualityGroup{
		Name:             model.DefaultGroupName,
		Moves:            a.Moves,
		States:           a.States,
		StateMoveToGroup: a.StateMoveToState,
		StateToGroup:     a.StateToSignal,
		GroupsCount:      countUniqueSignal(a),
	}
}

func EqualityGroupToMoore(a model.Moore, eg model.EqualityGroup) model.Moore {
	var newStates []string
	newStateToSignal := make(map[string]string)
	newStateMoveToState := make(map[string]string)
	checkedGroups := make(map[string]bool)

	for _, state := range eg.States {
		group := eg.StateToGroup[state]
		if _, found := checkedGroups[group]; found {
			continue
		}

		checkedGroups[group] = true
		newStates = append(newStates, state)
		newStateToSignal[state] = a.StateToSignal[state]

		for _, move := range a.Moves {
			stateMove := state + move

			resState := a.StateMoveToState[stateMove]

			newStateMoveToState[stateMove] = resState
		}
	}

	return model.Moore{
		Moves:            a.Moves,
		States:           newStates,
		StateToSignal:    newStateToSignal,
		StateMoveToState: newStateMoveToState,
	}
}

func removeUnreachableMooreStates(a *model.Moore) {
	reachedStates := make(map[string]bool)

	for _, state := range a.States {
		for _, move := range a.Moves {
			stateMove := state + move

			resState := a.StateMoveToState[stateMove]

			reachedStates[resState] = true
		}
	}

	a.States = getReachedState(a.States, reachedStates)
}

func countUniqueSignal(a model.Moore) int {
	uniqueSignal := make(map[string]bool)
	count := 0

	for _, state := range a.States {
		signal := a.StateToSignal[state]
		if _, found := uniqueSignal[signal]; !found {
			uniqueSignal[signal] = true
			count++
		}
	}

	return count
}
