package main

import (
	"os"
	"os/lw1/util"
	"os/lw2/mimization"
	"os/lw2/parse"
)

const minimizeMealyParam = "mealy"
const minimizeMooreParam = "moore"

func main() {
	records := util.ReadCSV(os.Args[2])

	var out [][]string
	switch os.Args[1] {
	case minimizeMealyParam:
		out = minimizeMealy(records)
		break
	case minimizeMooreParam:
		out = minimizeMoore(records)
	default:
		panic("undefiend arg")
	}

	util.WriteCSV(os.Args[3], out)
}

func minimizeMealy(records [][]string) [][]string {
	a := parse.TableToMealy(records)
	a = mimization.Mealy(a)

	return parse.MealyToTable(a)
}

func minimizeMoore(records [][]string) [][]string {
	return records
}
