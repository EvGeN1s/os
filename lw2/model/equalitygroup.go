package model

import "fmt"

const DefaultGroupName = 'a'

type EqualityGroup struct {
	Name             byte
	Moves            []string
	States           []string
	StateMoveToGroup map[string]string
	StateToGroup     map[string]string
	GroupsCount      int
}

func NextName(group EqualityGroup) byte {
	if group.Name += 1; group.Name > 'z' {
		return 'a'
	}
	return group.Name
}

func CreateGroupIndexName(name byte, i int) string {
	return fmt.Sprintf("%b%d", name, i)
}
