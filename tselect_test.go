package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetColumns(t *testing.T) {
	assert.Equal(t, []int{3, 5}, get_columns_to_select([]string{"%3", "%5"}))
	assert.Equal(t, []int{2, 1}, get_columns_to_select([]string{"%2", "%1"}))
	assert.Equal(t, []int{4, 5}, get_columns_to_select([]string{"%4", "%5"}))
}
