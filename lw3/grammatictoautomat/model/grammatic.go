package model

type Grammatic struct {
	States            []string
	Moves             []string
	StateMoveToStates map[string][]string
}
