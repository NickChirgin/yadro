package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nickchirgin/yadro/helpers"
)

func main() {
	data := helpers.Parser("data.csv")
	for _, val := range data {
		for _, v:= range val {
			if strings.HasPrefix(v, "=") {
				v = strconv.Itoa(helpers.Calculate(data, v))
			}
		}
	}
	fmt.Println(data)
}
