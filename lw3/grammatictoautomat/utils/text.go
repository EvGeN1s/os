package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func FileToRows(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	fContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}

	return strings.Split(string(fContent), "\n")
}
