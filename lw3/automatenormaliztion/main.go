package main

import (
	"fmt"
	"log"
	"os"
	"os/lw1/util"
	"os/lw3/automatenormaliztion/convert"
	"os/lw3/automatenormaliztion/parse"
	gramconvert "os/lw3/grammatictoautomat/convert"
	autmoateparse "os/lw3/grammatictoautomat/parse"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("invalid args count")
	}

	records := util.ReadCSV(os.Args[1])

	a := parse.TableToMultiAuto(records)

	gram := convert.RemoveEmptyMoves(a)
	fmt.Println(gram.States)
	fmt.Println(gram.StateMoveToStates)

	auto := gramconvert.GrammaticToAutomate(gram, 2)
	out := autmoateparse.AutomateToTable(auto)

	util.WriteCSV(os.Args[2], out)
}
