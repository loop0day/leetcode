package ex1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var target int
	var nums, want []int

	nums, target, want = []int{2, 7, 11, 15}, 9, []int{0, 1}
	assert.Equal(t, want, twoSum(nums, target))

	nums, target, want = []int{3, 2, 4}, 6, []int{1, 2}
	assert.Equal(t, want, twoSum(nums, target))

	nums, target, want = []int{3, 3}, 6, []int{0, 1}
	assert.Equal(t, want, twoSum(nums, target))

	nums, target, want = []int{3, 3, 3}, 6, []int{0, 1}
	assert.Equal(t, want, twoSum(nums, target))

	nums, target, want = []int{3, 3, 3}, 9, []int{}
	assert.Equal(t, want, twoSum(nums, target))
}
