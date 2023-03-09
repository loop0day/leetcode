package ex4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	var a, b []int
	var expected float64

	a, b, expected = []int{}, []int{}, 0
	assert.Equal(t, expected, findMedianSortedArrays(a, b))

	a, b, expected = []int{}, []int{1}, 1
	assert.Equal(t, expected, findMedianSortedArrays(a, b))

	a, b, expected = []int{1}, []int{}, 1
	assert.Equal(t, expected, findMedianSortedArrays(a, b))

	a, b, expected = []int{1, 3}, []int{2}, 2
	assert.Equal(t, expected, findMedianSortedArrays(a, b))

	a, b, expected = []int{2}, []int{1, 3}, 2
	assert.Equal(t, expected, findMedianSortedArrays(a, b))

	a, b, expected = []int{1, 3}, []int{2, 4}, 2.5
	assert.Equal(t, expected, findMedianSortedArrays(a, b))

	a, b, expected = []int{1, 2}, []int{3, 4, 5}, 3
	assert.Equal(t, expected, findMedianSortedArrays(a, b))
}
