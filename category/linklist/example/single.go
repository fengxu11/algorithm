package example

import (
	"bytes"
	"fmt"
)

// SingleLinkList 单链表
type SingleLinkList struct {
	dummyHead *Node
	size      int
}

// NewSingleLinkList 创建一个单链表
func NewSingleLinkList() *SingleLinkList {
	return &SingleLinkList{
		dummyHead: &Node{
			data: nil,
			next: nil,
		},
		size: 0,
	}
}

// GetSize 获取链表大小
func (s *SingleLinkList) GetSize() int {
	return s.size
}

// IsEmpty 链表是否为空
func (s *SingleLinkList) IsEmpty() bool {
	return s.size == 0
}

// Add 在index的位置插入元素
func (s *SingleLinkList) Add(index int, data interface{}) {

	if index < 0 || index > s.size {
		panic("Add error. index error")
	}

	prev := s.dummyHead
	// 循环到 index-1的位置、因为要找到index位置 之前的元素
	// 这样循环结束后、prev 也就是index之前的元素已经找到了
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	// node := NewNode(data, nil)
	// node.next = prev.next
	// prev.next = node

	// 这行的写法 和 上面三行是等价的
	prev.next = NewNode(data, prev.next)
	s.size++
}

// AddFirst 在链表头部添加节点
func (s *SingleLinkList) AddFirst(data interface{}) {

	s.Add(0, data)
}

// AddLast 添加到链表末尾
func (s *SingleLinkList) AddLast(data interface{}) {
	s.Add(s.size, data)
}

// Get 获取元素
func (s *SingleLinkList) Get(index int) interface{} {

	if index < 0 || index > s.size {
		panic("Add error. index error")
	}

	current := s.dummyHead.next

	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.data
}

// Set 更新
func (s *SingleLinkList) Set(index int, data interface{}) {

	if index < 0 || index > s.size {
		panic("Add error. index error")
	}

	current := s.dummyHead.next

	for i := 0; i < index; i++ {
		current = current.next
	}

	current.data = data
}

// Del 删除元素
func (s *SingleLinkList) Del(index int) interface{} {

	if index < 0 || index > s.size {
		panic("Add error. index error")
	}

	prev := s.dummyHead

	for i := 0; i < index; i++ {
		prev = prev.next
	}

	delNode := prev.next
	prev.next = delNode.next
	delNode.next = nil

	s.size--

	return delNode.data
}

// RemoveLast 删除第一个元素
func (s *SingleLinkList) RemoveFirst() interface{} {

	return s.Del(0)
}

// RemoveLast 删除最后一个元素
func (s *SingleLinkList) RemoveLast() interface{} {

	return s.Del(s.size - 1)
}

// Contains 链表中是否有这个元素
func (s *SingleLinkList) Contains(data interface{}) bool {

	current := s.dummyHead.next
	for current != nil {
		if current.data == data {
			return true
		}
		current = current.next
	}
	return false
}

func (s *SingleLinkList) String() string {

	var buffer bytes.Buffer

	current := s.dummyHead.next
	for current != nil {
		buffer.WriteString(fmt.Sprintf("%v -> ", current))
		current = current.next
	}

	buffer.WriteString("nil \n ")

	return buffer.String()
}
