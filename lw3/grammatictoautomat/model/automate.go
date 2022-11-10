package model

type Automate struct {
	States           []string
	Moves            []string
	StateMoveToState map[string]string
}
