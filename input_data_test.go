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
	assert.Equal(t, []int{0, 24, 34, 40, 49, 64}, findWordBeginnings(rows[0]))
	assert.Equal(t, []int{0, 24, 34, 49, 51, 57, 64}, findWordBeginnings(rows[1]))
}

func TestCalculateCoordinates(t *testing.T) {
	assert.Equal(t, []int{0, 24, 34, 49, 64}, getColumnIndices(rows))
}

func TestSafeSubstr(t *testing.T) {
	assert.Equal(t, "cd", "abcdefg"[2:4])
	assert.Equal(t, "cd", safeSubstr("abcdefg", 2, 4))
	assert.Equal(t, "cde", safeSubstr("abcdefg", 2, 5))
	assert.Equal(t, "cdef", safeSubstr("abcdefg", 2, 6))
	assert.Equal(t, "cdefg", safeSubstr("abcdefg", 2, 7))
	assert.Equal(t, "cdefg", safeSubstr("abcdefg", 2, 8))
	assert.Equal(t, "bcdefg", safeSubstr("abcdefg", 1, 8))
	assert.Equal(t, "abcdefg", safeSubstr("abcdefg", 0, 8))
	assert.Equal(t, "cdefg", safeSubstr("abcdefg", 2))
	assert.Equal(t, "defg", safeSubstr("abcdefg", 3))
	assert.Equal(t, "efg", safeSubstr("abcdefg", 4))
	assert.Equal(t, "fg", safeSubstr("abcdefg", 5))
	assert.Equal(t, "g", safeSubstr("abcdefg", 6))
	assert.Equal(t, "", safeSubstr("abcdefg", 7))
	assert.Equal(t, "", safeSubstr("abcdefg", 8))
}
