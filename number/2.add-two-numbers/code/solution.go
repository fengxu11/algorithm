package code

import "fmt"

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// ShowNode 遍历
func ShowNode(p *ListNode) {
	for p != nil {
		fmt.Println(*p)
		p = p.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// 判断两个单链表 是否为空
	if l1 == nil || l2 == nil {
		return nil
	}

	// 创建节点存储数据
	node1 := l1
	node2 := l2
	l3 := new(ListNode)
	node3 := l3

	// 标记
	tag := 0

	for node1 != nil || node2 != nil || tag > 0 {

		node3.Next = new(ListNode)
		node3 = node3.Next

		// 备份节点的value
		val1 := 0
		val2 := 0

		if node1 != nil {
			val1 = node1.Val
			node1 = node1.Next // 移动指针、遍历下一个节点
		}

		if node2 != nil {
			val2 = node2.Val
			node2 = node2.Next // 移动指针、遍历下一个节点
		}

		// 计算出node3 的value、也就是每次遍历的node1 + node2 + tag的值、 求余数
		node3.Val = (tag + val1 + val2) % 10

		// 计算出tag 也就是每次遍历的node1 + node2 + tag的值、 除以10 求值、 也就是 val1 + val2 大于 10的时候、则进一位
		tag = (tag + val1 + val2) / 10

		fmt.Println("val1: ", val1)
		fmt.Println("val2: ", val2)
		fmt.Println("tag: ", tag)
		fmt.Println("node3.Val: ", node3.Val)
		fmt.Println("")
		fmt.Println("")

	}

	return l3.Next
}
