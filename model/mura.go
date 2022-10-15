package model

type Mura []map[MuraState]int

type MuraState struct {
	State  int
	Signal int
}
