package model

type MultiStateAuto struct {
	States           []string
	Moves            []string
	FinalStates      map[string]bool
	StateMoveToState map[string][]string
}
