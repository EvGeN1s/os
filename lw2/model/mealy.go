package model

type Mealy struct {
	Moves                  []string
	States                 []string
	StateMoveToStateSignal map[string]string
}
