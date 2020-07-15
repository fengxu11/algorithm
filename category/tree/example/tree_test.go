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

	bst.preOrderNR()
}
