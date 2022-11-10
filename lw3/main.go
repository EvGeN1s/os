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
		gram = parse.RowsToLeftGrammatic(rows)
		break
	case transformRightSideGrammatic:
		gram = parse.RowsToRightGrammatic(rows)
		break
	default:
		log.Fatal("undefined arg")
	}

	auto := convert.GrammaticToAutomate(gram)
	out := parse.AutomateToTable(auto)

	util.WriteCSV(os.Args[3], out)
}

/* Test app
rows := []string{
"S -> aB",
"B -> aC | a | dB",
"C -> aB",
}

rows = utils.FileToRows("input-left.txt")

gram := parse.RowsToLeftGrammatic(rows)

auto := convert.GrammaticToAutomate(gram)

table := parse.AutomateToTable(auto)
csvutil.WriteCSV("output.csv", table)

fmt.Println(table)*/

/* Test model
gram := model.Grammatic{
States:            []string{"s", "b", "c", "f"},
Moves:             []string{"a", "d"},
StateMoveToStates: map[string][]string{
"sa": {"b"},
"ba": {"c", "f"},
"bd": {"b"},
"ca": {"b"},
},
}*/

/* list experiments
l := list.New()


l.PushBack("Hello")
l.PushBack(" world")

curr := l.Front()
str := fmt.Sprintf("%v", curr.Value)
fmt.Println(str)
l.Remove(curr)

curr = l.Front()
str = fmt.Sprintf("%v", curr.Value)
fmt.Println(str)
l.Remove(curr)*/
