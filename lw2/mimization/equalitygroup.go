package mimization

import (
	"os/lw2/model"
)

func EqualityGroup(eg model.EqualityGroup) model.EqualityGroup {
	newName := model.NextName(eg)
	newStateToGroup, count := findNewGroups(eg, newName)

	if eg.GroupsCount == count {
		return eg
	}

	return EqualityGroup(model.EqualityGroup{
		Name:             newName,
		Moves:            eg.Moves,
		States:           eg.States,
		StateMoveToGroup: createStateToGroupMap(eg, newStateToGroup),
		StateToGroup:     newStateToGroup,
		GroupsCount:      count,
	})
}

func findNewGroups(eg model.EqualityGroup, name byte) (map[string]string, int) {
	newStateToGroup := make(map[string]string)
	stateToGroup := make(map[string]string)
	count := 0

	for _, state := range eg.States {
		var states string
		for _, move := range eg.Moves {
			states += eg.StateMoveToGroup[state+move]
		}
		prevGroup := eg.StateToGroup[state]

		group, found := stateToGroup[prevGroup+states]
		if !found {
			group = model.CreateGroupIndexName(name, count)
			stateToGroup[prevGroup+states] = group
			count++
		}
		newStateToGroup[state] = group
	}

	return newStateToGroup, count
}

func createStateToGroupMap(eg model.EqualityGroup, stateToGroup map[string]string) map[string]string {
	stateMoveToGroup := make(map[string]string)

	for _, state := range eg.States {
		for _, move := range eg.Moves {
			stateMove := state + move

			resState := eg.StateMoveToGroup[state+move]
			group := stateToGroup[resState]

			stateMoveToGroup[stateMove] = group
		}
	}

	return stateMoveToGroup
}
