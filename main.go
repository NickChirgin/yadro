package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	rec, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for i, val := range rec {
		for j, v := range val {

			if strings.HasPrefix(v, "=") {
				fmt.Println(v, i, j)
			}
		}
	}
}
