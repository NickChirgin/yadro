package helpers

import (
	"strconv"
	"strings"
)

func Calculate(csv [][]string, cell string) int {
	cells, math := StringToCell(csv, cell)	
	if len(cells) == 1 {
		if strings.HasPrefix(cells[0], "=") {
			return Calculate(csv, cells[0][1:])
		}
		n, _ := strconv.Atoi(cells[0])
		return n 
	}
	first, second := cells[0], cells[1]
	if strings.HasPrefix(first, "=") {
		first = strconv.Itoa(Calculate(csv, first[1:]))
	}
	if strings.HasPrefix(second, "=") {
		second = strconv.Itoa(Calculate(csv, second[1:]))
	}
	num1, _ := strconv.Atoi(first)
	num2, _ := strconv.Atoi(second)
	switch math {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 == 0 {
			return 0
		}	else {
			return num1/num2
		}
	}
	return num1
}