package ex2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func build(nums []int) *ListNode {
	l := &ListNode{}
	t := l

	for _, num := range nums {
		t.Next = &ListNode{Val: num, Next: nil}
		t = t.Next
	}

	return l.Next
}

func equal(l *ListNode, r *ListNode) bool {
	for (l != nil) && (r != nil) {
		if l.Val != r.Val {
			return false
		}

		l = l.Next
		r = r.Next
	}

	return (l == nil) && (r == nil)
}

func TestAddTwoNumbers(t *testing.T) {
	var l1, l2, expected, actual *ListNode

	l1, l2 = build([]int{2, 4, 3}), build([]int{5, 6, 4})
	expected, actual = build([]int{7, 0, 8}), addTwoNumbers(l1, l2)
	assert.True(t, equal(expected, actual))

	l1, l2 = build([]int{0}), build([]int{0})
	expected, actual = build([]int{0}), addTwoNumbers(l1, l2)
	assert.True(t, equal(expected, actual))

	l1, l2 = build([]int{9, 9, 9, 9, 9, 9, 9}), build([]int{9, 9, 9, 9})
	expected, actual = build([]int{8, 9, 9, 9, 0, 0, 0, 1}), addTwoNumbers(l1, l2)
	assert.True(t, equal(expected, actual))

	l1, l2 = build([]int{9, 9, 9, 9}), build([]int{9, 9, 9, 9, 9, 9, 9})
	expected, actual = build([]int{8, 9, 9, 9, 0, 0, 0, 1}), addTwoNumbers(l1, l2)
	assert.True(t, equal(expected, actual))
}
