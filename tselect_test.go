package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetingOneColumn(t *testing.T) {
	result, _ := get_columns_to_select([]string{"%5"})
	assert.Equal(t, []int{5}, result)
}

func TestGetColumns(t *testing.T) {
	result, _ := get_columns_to_select([]string{"%3", "%5"})
	assert.Equal(t, []int{3, 5}, result)
}

func TestGetColumnsInReverseOrder(t *testing.T) {
	result, _ := get_columns_to_select([]string{"%2", "%1"})
	assert.Equal(t, []int{2, 1}, result)
}

func TestGettingColumnZero(t *testing.T) {
	_, err := get_columns_to_select([]string{"%0"})
	assert.NotNil(t, err)
}

func TestGettingNegativeColumn(t *testing.T) {
	_, err := get_columns_to_select([]string{"%-1"})
	assert.NotNil(t, err)
}
