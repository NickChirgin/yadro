package helpers

func Mapper(table *Excel) (map[string]int, map[string]int) {
	data := table.Table
	columnMap := map[string]int{} 
	rowMap := map[string]int{} 
	for i, v:= range data[0] {
		columnMap[v] = i
	}
	for j, val := range data {
		rowMap[val[0]] = j
	}
	return rowMap, columnMap
}