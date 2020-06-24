### No. 2
### Title: add-two-number  两数相加
### Tag: array
### Description: 
```

给出两个 非空 的链表、 用来表示两个非负的整数、 其中它们各自的位数 是按照 逆序的方式存储的、
并且 它们的每个节点只能存储一位数字

如果、我们将这两个数加起来、则会返回一个新的链表来表示它们的和
您也可以假设 除了数字0 之外、这两个数都不会以 0开头


```

### example
```

示例 1:
    输入: (2 -> 4 -> 3) + (5 -> 6 ->  4)
    输出: 7 -> 0 -> 8
    原因: 342 + 465 = 807


```


### 解题思路

```

读题得知: 
    1. 两个链表链表都是 非空的、 所以至少是 有一个节点的、
    2. 链表存储的是非负整数、且都是按照逆序的形式存储的、例如 342  存储为 2 -> 4 -> 3
    3. 非负整数不会以 0开头、因此不需要考虑链表末尾有 无数个0的情况
    4. *** 两数相加后、大于等于10、则需要进位
    5. 最后两数相加的结果 也要存储为链表返回、并且按照逆序表示


```

### Code
```go

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

```