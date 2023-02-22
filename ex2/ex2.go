package ex2

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s, c := 0, 0
	l := &ListNode{}
	t := l

	for (l1 != nil) && (l2 != nil) {
		s = l1.Val + l2.Val + c
		c = s / 10
		s = s % 10

		t.Next = &ListNode{Val: s, Next: nil}
		t = t.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	if l2 != nil {
		l1 = l2
	}

	for l1 != nil {
		s = l1.Val + c
		c = s / 10
		s = s % 10

		t.Next = &ListNode{Val: s, Next: nil}
		t = t.Next

		l1 = l1.Next
	}

	if c != 0 {
		t.Next = &ListNode{Val: c, Next: nil}
	}

	return l.Next
}
