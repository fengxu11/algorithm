package example

import (
	"testing"
)

func TestNewBST(t *testing.T) {

	// input
	/////////////////////////////////
	//				5			   //
	//			/		\		   //
	//		   3		 6		   //
	//		  /	\		  \		   //
	//       2   4         8	   //
	/////////////////////////////////

	// 创建一个 二分搜索树、并添加数据
	bst := NewBST()

	nums := []int{5, 3, 6, 8, 4, 2}

	for _, num := range nums {

		bst.Add(num)
	}

	// 前中后序 遍历二分搜索树
	//bst.PreOrder()

	//bst.InOrder()

	//bst.PostOrder()

	// 非递归(循环+栈) 前序遍历
	//bst.preOrderNR()

	// 层序遍历 循环 + 队列(或链表)
	bst.levelOrder()
}
