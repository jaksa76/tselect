package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type InputData struct {
	in                 *bufio.Scanner
	column_coordinates []int
	rows               []string
	current_row        int
	text_from_scanner  string
}

func read_some_rows(in *bufio.Scanner, n int) []string {
	s := make([]string, 0, n)
	for i := 0; i < n && in.Scan(); i++ {
		s = append(s, in.Text())
	}
	return s
}

func intersect(a, b []int) []int {
	result := make([]int, 0)
	for ia, ib := 0, 0; ia < len(a) && ib < len(b); {
		if a[ia] == b[ib] {
			result = append(result, a[ia])
			ia++
			ib++
		} else if a[ia] > b[ib] {
			ib++
		} else if a[ia] < b[ib] {
			ia++
		}
	}
	return result
}

func find_word_beginnings(s string) []int {
	regex, _ := regexp.Compile(`\s\w`)
	newCoords := make([]int, 1)
	newCoords[0] = 0
	for _, coord := range regex.FindAllStringIndex(s, -1) {
		newCoords = append(newCoords, coord[0]+1)
	}
	return newCoords
}

func calculate_coordinates(initial_rows []string) []int {
	coords := find_word_beginnings(initial_rows[0])
	for i, row := range initial_rows {
		if i > 0 {
			newCoords := find_word_beginnings(row)
			coords = intersect(coords, newCoords)
		}
	}
	return coords
}

func NewInputData(in *os.File) *InputData {
	scanner := bufio.NewScanner(in)
	initial_rows := read_some_rows(scanner, 10)
	return &InputData{scanner, calculate_coordinates(initial_rows), initial_rows, 0, ""}
}

func (i *InputData) Scan() bool {
	if i.current_row < len(i.rows) {
		i.current_row++
		return true
	} else {
		hasMore := i.in.Scan()
		if hasMore {
			i.text_from_scanner = i.in.Text()
		}
		return hasMore
	}
}

func (i *InputData) get_curent_row() string {
	row := i.text_from_scanner
	if i.current_row < len(i.rows) {
		row = i.rows[i.current_row]
	}
	return row
}

func safe_substr(s string, bounds ...int) string {
	if len(bounds) < 1 || len(bounds) > 2 {
		panic(fmt.Errorf("invalid bounds"))
	}

	start := bounds[0]
	if start > len(s) {
		return ""
	}

	if len(bounds) == 1 {
		return s[start:]
	}

	end := bounds[1]
	if end > len(s) {
		return s[start:]
	}

	return s[start:end]
}

func (i *InputData) GetColumn(n int) string {
	n = n - 1
	row := i.get_curent_row()

	lastCol := len(i.column_coordinates) - 1
	if n > lastCol {
		return ""
	}

	start := i.column_coordinates[n]
	if n == lastCol {
		return safe_substr(row, start)
	}

	end := i.column_coordinates[n+1]
	return safe_substr(row, start, end)
}

func (i *InputData) IsEmptyRow() bool {
	return len(i.get_curent_row()) == 0
}
