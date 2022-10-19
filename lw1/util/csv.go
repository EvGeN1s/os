package util

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func ReadCSV(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func WriteCSV(filePath string, data [][]string) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	_, _ = f.Seek(0, io.SeekStart)
	_ = f.Truncate(0)
	defer f.Close()

	csvWriter := csv.NewWriter(f)
	csvWriter.Comma = ';'
	err = csvWriter.WriteAll(data)
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
}
