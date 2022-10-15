package model

type Milli []map[int]MilliState

type MilliState struct {
	State  int
	Signal int
}
