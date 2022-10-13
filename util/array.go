package util

import "os/model"

func ContainsMuraState(s []model.MuraState, e model.MuraState) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsMilliState(s []model.MilliState, e model.MilliState) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
