package convert

import (
	"os/lw2/model"
	"strings"
)

type stateWithSignal struct {
	State  string
	Signal string
}

func MealyToEqualityGroup(a model.Mealy) model.EqualityGroup {
	removeUnreachableMealyStates(&a)
	stateToGroup, count := createMealyStateToEqualityGroup(a)

	return model.EqualityGroup{
		Name:             model.DefaultGroupName,
		Moves:            a.Moves,
		States:           a.States,
		StateMoveToGroup: createMealyStateMoveToGroupMap(a, stateToGroup),
		StateToGroup:     stateToGroup,
		GroupsCount:      count,
	}
}

func EqualityGroupToMealy(a model.Mealy, eg model.EqualityGroup) model.Mealy {
	var newStates []string
	newStateMoveToState := make(map[string]string)
	checkedGroups := make(map[string]bool)

	for _, state := range eg.States {
		group := eg.StateToGroup[state]
		if _, found := checkedGroups[group]; found {
			continue
		}

		checkedGroups[group] = true
		newStates = append(newStates, state)

		for _, move := range a.Moves {
			stateMove := state + move

			resState := a.StateMoveToStateSignal[stateMove]

			newStateMoveToState[stateMove] = resState
		}
	}

	return model.Mealy{
		Moves:                  a.Moves,
		States:                 newStates,
		StateMoveToStateSignal: newStateMoveToState,
	}
}

func createMealyStateToEqualityGroup(a model.Mealy) (map[string]string, int) {
	stateToGroup := make(map[string]string)
	signalsToGroup := make(map[string]string)
	count := 0

	for _, state := range a.States {
		var signals string
		for _, move := range a.Moves {
			stateWithSignal := getStateWithSignalFromString(a.StateMoveToStateSignal[state+move])
			signals += stateWithSignal.Signal
		}
		group, found := signalsToGroup[signals]
		if !found {
			group = model.CreateGroupIndexName(model.DefaultGroupName, count)
			count++
			signalsToGroup[signals] = group
		}
		stateToGroup[state] = group
	}

	return stateToGroup, count
}

func createMealyStateMoveToGroupMap(a model.Mealy, stateToGroup map[string]string) map[string]string {
	stateMoveToGroup := make(map[string]string)

	for _, state := range a.States {
		for _, move := range a.Moves {
			stateMove := state + move

			stateWithSignal := getStateWithSignalFromString(a.StateMoveToStateSignal[stateMove])
			resState := stateWithSignal.State
			group := stateToGroup[resState]

			stateMoveToGroup[stateMove] = group
		}
	}

	return stateMoveToGroup
}

func removeUnreachableMealyStates(a *model.Mealy) {
	reachedStates := make(map[string]bool)

	for _, state := range a.States {
		for _, move := range a.Moves {
			stateMove := state + move

			stateWithSignal := getStateWithSignalFromString(a.StateMoveToStateSignal[stateMove])
			resState := stateWithSignal.State

			reachedStates[resState] = true
		}
	}

	a.States = getReachedState(a.States, reachedStates)
}

func getReachedState(states []string, reachedStates map[string]bool) []string {
	var resStates []string

	for _, state := range states {
		if _, found := reachedStates[state]; found {
			resStates = append(resStates, state)
		}
	}

	return resStates
}

func getStateWithSignalFromString(s string) stateWithSignal {
	res := strings.Split(s, "/")
	return stateWithSignal{
		State:  res[0],
		Signal: res[1],
	}
}
