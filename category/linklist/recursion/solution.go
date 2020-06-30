package recursion

func removeElementsByFor(head *ListNode, val int) *ListNode {

	for head != nil && head.Val == val {
		// delNode := head
		// head = head.Next
		// delNode.Next = nil
		head = head.Next
	}

	if head == nil {
		return nil
	}

	prev := head
	for prev.Next != nil {
		// 如果链表当前节点中的value 和  val相等、就删除
		if prev.Next.Val == val {
			// delNode := prev.Next
			// prev.Next = delNode.Next
			// delNode.Next = nil
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}

	}

	return head
}

func removeElementsByDummyHead(head *ListNode, val int) *ListNode {

	dummyHead := &ListNode{-1, nil}
	dummyHead.Next = head

	prev := dummyHead
	for prev.Next != nil {
		// 如果链表当前节点中的value 和  val相等、就删除
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}

	}

	return dummyHead.Next
}
