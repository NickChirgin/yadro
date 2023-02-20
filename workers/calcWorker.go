package workers

import "github.com/nickchirgin/yadro/helpers"
func CalcWorker(table *helpers.Excel, tasks <-chan string) {
	for task := range tasks {
		table.Table[i][j] = helpers.Calculate(table, task)
	}
} 