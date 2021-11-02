package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var colSelector = regexp.MustCompile(`^\%\d+$`)

func get_columns_to_select(args []string) ([]int, error) {
	columns := make([]int, 0, len(args))
	for _, arg := range args {
		if colSelector.MatchString(arg) {
			n, err := strconv.Atoi(arg[1:])
			if err != nil {
				return nil, fmt.Errorf("Could not parse %s", arg)
			} else if n == 0 {
				return nil, fmt.Errorf("Please use 1-based indices for columns")
			} else if n < 0 {
				return nil, fmt.Errorf("Negative index %d is not allowed", n)
			}
			columns = append(columns, n)
		} else {
			return nil, fmt.Errorf("Cannot determine column argument %s", arg)
		}
	}
	return columns, nil
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
	scanner := bufio.NewScanner(os.Stdin)

	if strings.HasPrefix(args[0], "-s") {
		linesToSkip, _ := strconv.Atoi(args[0][2:])
		for i := 0; i < linesToSkip; i++ {
			scanner.Scan()
		}
		args = args[1:]
	}
	columns, err := get_columns_to_select(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	inputData := NewInputData(scanner)
	for inputData.HasMoreLines() {
		if !inputData.IsEmptyRow() {
			values := select_columns(columns, inputData)
			fmt.Println(strings.Join(values, ""))
		}
	}
}
