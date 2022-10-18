package main

import (
	"os"
	convert2 "os/lw1/convert"
	parse2 "os/lw1/parse"
	util2 "os/lw1/util"
)

const milliToMure = "mealy-to-moore"
const mureToMilli = "moore-to-mealy"

func main() {
	records := util2.ReadCSV(os.Args[2])

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

	util2.WriteCSV(os.Args[3], out)
}

func convertMilli(records [][]string) [][]string {
	in := parse2.ConvertFromTableToMilli(records)
	out := convert2.MilliToMure(in)
	return parse2.ConvertFromMuraToTable(out)
}

func convertMure(records [][]string) [][]string {
	in := parse2.ConvertFromTableToMura(records)
	out := convert2.MuraToMilli(in)
	return parse2.ConvertFromMilliToTable(out)
}
