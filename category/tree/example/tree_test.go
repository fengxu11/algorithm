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
	//bst := NewBST()
	//
	//nums := []int{5, 3, 6, 8, 4, 2}
	//
	//for _, num := range nums {
	//
	//	bst.Add(num)
	//}

	// 前中后序 遍历二分搜索树
	//bst.PreOrder()

	//bst.InOrder()

	//bst.PostOrder()

	// 非递归(循环+栈) 前序遍历
	//bst.preOrderNR()

	// 层序遍历 循环 + 队列(或链表)
	//bst.levelOrder()


	//bst := NewBST()
	//
	//for i := 0; i < 1000; i++ {
	//	bst.Add(rand.Intn(10000))
	//}
	//
	//arr := make([]interface{}, 0)
	//
	//for !bst.IsEmpty() {
	//	arr = append(arr, bst.RemoveMin().Data)
	//}
	//
	//fmt.Println(arr)
	//
	//fmt.Println(len(arr))
	//
	//// 测试删除最小值
	//for i := 1; i < len(arr); i++ {
	//
	//	if arr[i-1].(int) > arr[i].(int) {
	//		panic("算法有问题")
	//	}
	//}
	//
	//fmt.Println("bst: ", bst)




}
