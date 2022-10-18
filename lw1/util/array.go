package util

import (
	model2 "os/lw1/model"
)

func ContainsState(s []model2.MilliState, e model2.MilliState) bool {
	for _, a := range s {
		if a.State == e.State && a.Signal == e.Signal {
			return true
		}
	}
	return false
}
