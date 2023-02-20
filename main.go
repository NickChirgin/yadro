package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nickchirgin/yadro/helpers"
)

func main() {
	filePath, _  := filepath.Abs(os.Args[1])
	fmt.Println(filePath)
	excel := helpers.Parser(filePath)
	data := excel.Table
	for i, val := range data {
		for j, v:= range val {
			v = strings.Trim(v, " ")
			if strings.HasPrefix(v, "=") {
				data[i][j] = helpers.Calculate(&excel, v[1:])
			}
		}
	}
	for _, val := range data {
		fmt.Println(strings.Join(val, ","))
	}

}
