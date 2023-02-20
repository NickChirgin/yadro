package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nickchirgin/yadro/helpers"
)

func main() {
	data := helpers.Parser("data.csv")

	for i, val := range data {
		for j, v:= range val {
			v = strings.Trim(v, " ")
			if strings.HasPrefix(v, "=") {
				result := helpers.Calculate(data, v[1:])
				data[i][j] = strconv.Itoa(result)
				v = strconv.Itoa(result)
			}
		}
	}
	for _, val := range data {
		fmt.Println(strings.Join(val, ","))
	}

}
