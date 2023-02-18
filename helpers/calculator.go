package helpers

import (
	"log"
	"regexp"
	"strconv"
)

func Calculate(csv [][]string, cell string) int {
	regexNum := regexp.MustCompile("[0-9]+")
	split := regexNum.Split(cell, -1)
	numbers := regexNum.FindAllStringSubmatch(cell, -1)
	rowMap, colMap := Mapper(csv)
	secondCol := split[1][1:]
	math := split[1][:1]
	secondNum := ""
	if secondCol == " " {
		secondNum = numbers[0][1]
	} else {
		secondNum = csv[rowMap[numbers[1][0]]][colMap[split[1][1:]]]
	}
	firstNum := csv[rowMap[numbers[0][0]]][colMap[split[0]]]
	num1, _ := strconv.Atoi(firstNum)
	num2, _ := strconv.Atoi(secondNum)
	switch math {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 == 0 {
			log.Fatal("Division by 0")
		}	else {
			return num1/num2
		}
	}
	return num1
}