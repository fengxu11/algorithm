package example

import (
	"bytes"
	"fmt"
)

// S 栈 接口
type S interface {
	GetSize() int
	IsEmpty() bool
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
}

// LinkedListStack 使用链表实现栈
type LinkedListStack struct {
	linkedList *SingleLinkList
}

// NewLinkedListStack 创建栈
func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		linkedList: &SingleLinkList{
			dummyHead: &Node{},
			size:      0,
		},
	}
}

// GetSize 获取栈的大小
func (s *LinkedListStack) GetSize() int {

	return s.linkedList.size
}

// IsEmpty 栈是否为空
func (s *LinkedListStack) IsEmpty() bool {
	return s.linkedList.IsEmpty()
}

// Push 入栈
func (s *LinkedListStack) Push(data interface{}) {
	s.linkedList.AddFirst(data)
}

// Pop 出栈
func (s *LinkedListStack) Pop() interface{} {
	return s.linkedList.RemoveFirst()
}

// Peek 查看第一个元素
func (s *LinkedListStack) Peek() interface{} {
	return s.linkedList.Get(0)
}

func (s *LinkedListStack) String() string {

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Stack: top %s", s.linkedList))
	return buffer.String()
}
