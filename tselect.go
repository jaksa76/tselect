package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var colSelector = regexp.MustCompile(`^\%\d+$`)

func get_columns_to_select(args []string) []int {
	columns := make([]int, 0, len(args))
	for _, arg := range args {
		if colSelector.MatchString(arg) {
			n, _ := strconv.Atoi(arg[1:])
			columns = append(columns, n)
		}
	}
	return columns
}

func select_columns(columns []int, text *InputData) []string {
	result := make([]string, 0)
	for _, c := range columns {
		result = append(result, text.GetColumn(c))
	}
	return result
}

func main() {
	args := os.Args[1:]
	columns := get_columns_to_select(args)

	inputData := NewInputData(os.Stdin)
	for inputData.Scan() {
		if !inputData.IsEmptyRow() {
			values := select_columns(columns, inputData)
			fmt.Println(strings.Join(values, ""))
		}
	}
}
