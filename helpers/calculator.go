package helpers

import (
	"log"
	"strconv"
	"strings"
)

func Calculate(csv *Excel, cell string) string {
	cells, math := StringToCell(csv, cell)	
	if math == "" {
		if strings.HasPrefix(cells[0], "=") {
			return Calculate(csv, cells[0][1:])
		}
		return cells[0] 
	}
	first, second := cells[0], cells[1]
	if strings.HasPrefix(first, "=") {
		first = Calculate(csv, first[1:])
	}
	if strings.HasPrefix(second, "=") {
		second = Calculate(csv, second[1:])
	}
	num1, _ := strconv.Atoi(first)
	num2, _ := strconv.Atoi(second)
	switch math {
	case "+":
		return strconv.Itoa(num1 + num2)
	case "-":
		return strconv.Itoa(num1 - num2)
	case "*":
		return strconv.Itoa(num1 * num2)
	case "/":
		if num2 == 0 {
			log.Fatalf("Division by zero in %s", cell)
		}	else {
			return strconv.Itoa(num1/num2)
		}
	}
	return strconv.Itoa(num1)
}