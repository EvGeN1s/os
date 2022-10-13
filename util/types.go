package util

import (
	"regexp"
	"strconv"
)

func StringToInt(s string) int {
	re := regexp.MustCompile("[0-9]+")
	numS := re.FindString(s)
	res, err := strconv.Atoi(numS)
	if err != nil {
		panic(err)
	}

	return res
}
