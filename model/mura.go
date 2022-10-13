package model

type Mura map[int]map[MuraState]int

type MuraState struct {
	State  int
	Signal int
}
