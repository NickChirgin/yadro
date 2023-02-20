package helpers

import (
	"encoding/csv"
	"log"
	"os"
)
type Excel struct {
	Table [][]string
}
func Parser(filePath string) Excel{
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to open file %s", filePath)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Cant read csv file %s", filePath)
	}

	return	Excel{Table: data} 
}