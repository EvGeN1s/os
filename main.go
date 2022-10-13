package main

import (
	"fmt"
	"os"
	"os/convert"
	"os/parse"
	"os/util"
)

func main() {
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])

	records := util.ReadCSV(os.Args[2])
	in := parse.ConvertFromTableToMilli(records)
	out := convert.ConvertMilliToMure(in)
	res := parse.ConvertMuraToOutStr(out)

	fmt.Println("Go")
}
