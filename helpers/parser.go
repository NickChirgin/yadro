package helpers

import (
	"encoding/csv"
	"log"
	"os"
)

func Parser(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to open file")
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Cant read csv file")
	}

	return data
}