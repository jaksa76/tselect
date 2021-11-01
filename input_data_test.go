package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var rows []string = []string{
	"REPOSITORY              TAG       IMAGE ID       CREATED        SIZE",
	"python                  latest    c05c608cfa20   2 weeks ago    915MB",
	"bkimminich/juice-shop   latest    70ef66aed903   2 weeks ago    494MB",
	"debian                  latest    f776cfb21b5e   2 weeks ago    124MB",
	"alpine                  latest    14119a10abf4   2 months ago   5.6MB",
}

func TestIntersection(t *testing.T) {
	assert.Equal(t, []int{1}, intersect([]int{1}, []int{1}))
	assert.Equal(t, []int{3}, intersect([]int{1, 3}, []int{2, 3}))
	assert.Equal(t, []int{3}, intersect([]int{1, 3}, []int{2, 3}))
	assert.Equal(t, []int{2, 3}, intersect([]int{1, 2, 3}, []int{2, 3, 4}))
	assert.Equal(t, []int{0, 24, 34, 49, 64}, intersect([]int{0, 24, 34, 40, 49, 64}, []int{0, 24, 34, 49, 51, 57, 64}))
}

func TestFindWordBegninnings(t *testing.T) {
	assert.Equal(t, []int{0, 24, 34, 40, 49, 64}, find_word_beginnings(rows[0]))
	assert.Equal(t, []int{0, 24, 34, 49, 51, 57, 64}, find_word_beginnings(rows[1]))
}

func TestCalculateCoordinates(t *testing.T) {
	assert.Equal(t, []int{0, 24, 34, 49, 64}, calculate_coordinates(rows))
}
