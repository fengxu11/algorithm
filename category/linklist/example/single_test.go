package example

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {

	singleLinkList := NewSingleLinkList()

	fmt.Println("单链表是否为空: ", singleLinkList.IsEmpty())

	for i := 0; i < 5; i++ {
		singleLinkList.AddFirst(i)
		fmt.Println(singleLinkList)
	}

	singleLinkList.Add(2, 666)
	fmt.Println(singleLinkList)

	fmt.Println("666这个元素是否存在: ", singleLinkList.Contains(666))

	fmt.Println("获取 index 为1 位置上的元素: ", singleLinkList.Get(1))

	singleLinkList.Set(1, 333)
	fmt.Print("修改 index 为1 位置上的元素改为 333: ")
	fmt.Println(singleLinkList)

	singleLinkList.AddLast(99)
	fmt.Println(singleLinkList)

	fmt.Println("单链表大小: ", singleLinkList.GetSize())

	fmt.Println("单链表是否为空: ", singleLinkList.IsEmpty())

	fmt.Println("")

	fmt.Println("删除位置为1 的元素: ", singleLinkList.Del(1))
	fmt.Println(singleLinkList)

}
