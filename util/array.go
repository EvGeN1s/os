package util

import (
	"os/model"
)

func ContainsState(s []model.MilliState, e model.MilliState) bool {
	for _, a := range s {
		if a.State == e.State && a.Signal == e.Signal {
			return true
		}
	}
	return false
}
