package helpers

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func StringToCell(csv [][]string, cell string) ([]string, string) {
	regexNum := regexp.MustCompile("[0-9]+")
	split := regexNum.Split(cell, -1)
	rowMap, colMap := Mapper(csv)
	numbers := regexNum.FindAllStringSubmatch(cell, -1)
	numSlice := []string{}
	math := "" 
	for i, val := range numbers {
		if strings.Trim(split[i], " ") == "" {
			numSlice = append(numSlice, val[0])
		} else {
			col := split[i]
			if i == 1 {
				col = split[i][1:]
				math = split[i][:1]
			}
			if _, ok := colMap[col]; !ok {
				log.Fatal(fmt.Errorf("Column %s doesn't exist", col))
			}
			if _, ok := rowMap[val[0]]; !ok{
				log.Fatal(fmt.Errorf("Row %s doesn't exist", val[0]))
			}
			numSlice = append(numSlice, csv[rowMap[val[0]]][colMap[col]])
		}
	}
	return numSlice, math 
}