package main

import (
	"log"
	"os"
	"os/lw1/util"
	"os/lw3/grammatictoautomat/convert"
	"os/lw3/grammatictoautomat/model"
	"os/lw3/grammatictoautomat/parse"
	"os/lw3/grammatictoautomat/utils"
)

const (
	transformLeftSideGrammatic  = "left"
	transformRightSideGrammatic = "right"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("invalid args count")
	}

	rows := utils.FileToRows(os.Args[2])

	var gram model.Grammatic
	switch os.Args[1] {
	case transformLeftSideGrammatic:
		gram = parse.RowsToLGrammatic(rows)
		break
	case transformRightSideGrammatic:
		gram = parse.RowsToRGrammatic(rows)
		break
	default:
		log.Fatal("undefined arg")
	}

	auto := convert.GrammaticToAutomate(gram, 1)
	out := parse.AutomateToTable(auto)

	util.WriteCSV(os.Args[3], out)
}
