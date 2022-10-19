package model

type Moore struct {
	Moves            []string
	States           []string
	StateToSignal    map[string]string
	MoveStateToState map[string]string
}
