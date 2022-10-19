package main

import (
	"log"
	"os"
	"os/lw1/util"
	"os/lw2/mimization"
	"os/lw2/parse"
)

const minimizeMealyParam = "mealy"
const minimizeMooreParam = "moore"

func main() {
	if len(os.Args) < 3 {
		log.Fatal("invalid args count")
	}

	records := util.ReadCSV(os.Args[2])

	var out [][]string
	switch os.Args[1] {
	case minimizeMealyParam:
		out = minimizeMealy(records)
		break
	case minimizeMooreParam:
		out = minimizeMoore(records)
	default:
		log.Fatal("undefined arg")
	}

	util.WriteCSV(os.Args[3], out)
}

func minimizeMealy(records [][]string) [][]string {
	a := parse.TableToMealy(records)
	a = mimization.Mealy(a)

	return parse.MealyToTable(a)
}

func minimizeMoore(records [][]string) [][]string {
	a := parse.TableToMoore(records)
	a = mimization.Moore(a)

	return parse.MooreToTable(a)
}
