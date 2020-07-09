package recursion

func removeElements(head *ListNode, val int) *ListNode {

	if head == nil {
		return nil
	}

	// res := removeElements(head.Next, val)
	// if head.Val == val {
	// 	return res
	// }

	// head.Next = res
	// return head

	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}

	return head
}
