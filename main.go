package main

import (
	"os"
	"os/convert"
	"os/parse"
	"os/util"
)

const milliToMure = "mealy-to-moore"
const mureToMilli = "moore-to-mealy"

func main() {
	records := util.ReadCSV(os.Args[2])

	var out [][]string
	switch os.Args[1] {
	case milliToMure:
		out = convertMilli(records)
		break
	case mureToMilli:
		out = convertMure(records)
	default:
		panic("undefiend arg")
	}

	util.WriteCSV(os.Args[3], out)
}

func convertMilli(records [][]string) [][]string {
	in := parse.ConvertFromTableToMilli(records)
	out := convert.MilliToMure(in)
	return parse.ConvertFromMuraToTable(out)
}

func convertMure(records [][]string) [][]string {
	in := parse.ConvertFromTableToMura(records)
	out := convert.MuraToMilli(in)
	return parse.ConvertFromMilliToTable(out)
}
