package ex3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var s string
	var expected int

	s, expected = "abcabcbb", 3
	assert.Equal(t, expected, lengthOfLongestSubstring(s))

	s, expected = "bbbbb", 1
	assert.Equal(t, expected, lengthOfLongestSubstring(s))

	s, expected = "pwwkew", 3
	assert.Equal(t, expected, lengthOfLongestSubstring(s))

	s, expected = "", 0
	assert.Equal(t, expected, lengthOfLongestSubstring(s))

	s, expected = "a", 1
	assert.Equal(t, expected, lengthOfLongestSubstring(s))

	s, expected = "abc", 3
	assert.Equal(t, expected, lengthOfLongestSubstring(s))
}
