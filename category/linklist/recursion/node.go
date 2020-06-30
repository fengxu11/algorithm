package recursion

import (
	"bytes"
	"fmt"
)

// ListNode leetcode 203题
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNodeByArr 根据一个数组创建一个链表
func NewListNodeByArr(arr []int) *ListNode {
	if len(arr) == 0 {
		panic("param error.")
	}

	head := &ListNode{
		Val: arr[0],
	}

	current := head

	for i := 1; i < len(arr); i++ {
		current.Next = &ListNode{Val: arr[i]}
		current = current.Next
	}

	return head
}

func (s *ListNode) String() string {

	var buffer bytes.Buffer

	current := s
	for current != nil {
		buffer.WriteString(fmt.Sprintf("%v ->", current.Val))
		current = current.Next
	}

	buffer.WriteString(" nil")

	return buffer.String()
}
