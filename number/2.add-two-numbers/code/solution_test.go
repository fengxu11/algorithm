package code

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

	node1 := new(ListNode)
	node1.Val = 2
	node1.Next = &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  3,
			Next: nil,
		},
	}

	node2 := new(ListNode)
	node2.Val = 5
	node2.Next = &ListNode{
		Val: 6,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}

	ShowNode(addTwoNumbers(node1, node2))
}
