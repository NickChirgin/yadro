package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/nickchirgin/yadro/helpers"
)

func main() {
	filePath, _  := filepath.Abs(os.Args[1])
	var wg sync.WaitGroup
	excel := helpers.Parser(filePath)
	data := excel.Table
	for i, val := range data {
		for j, v:= range val {
			v = strings.Trim(v, " ")
			if strings.HasPrefix(v, "=") {
				wg.Add(1)
				go func(i, j int, v string){
					data[i][j] = helpers.Calculate(&excel, v[1:])
					wg.Done()
				}(i, j, v)
			}
		}
	}
	wg.Wait()
	for _, val := range excel.Table { 
		fmt.Println(strings.Join(val, ","))
	}
}
