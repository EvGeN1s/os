package model

type Milli map[int]map[int]MilliState

type MilliState struct {
	State  int
	Signal int
}
