package example

import (
	"bytes"
	"fmt"
)

// 链表队列接口
type IQueue interface {
	GetSize() int
	IsEmpty() bool
	Enqueue(e interface{})
	Dequeue() interface{}
	GetFront() interface{}
}

// LinkedListQueue 链表实现的队列
type LinkedListQueue struct {
	head *Node
	tail *Node
	size int
}

// NewLinkListQueue 创建一个链表队列
func NewLinkListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// GetSize 获取链表队列 size
func (q *LinkedListQueue) GetSize() int {
	return q.size
}

// IsEmpty 判断链表队列是否为空
func (q *LinkedListQueue) IsEmpty() bool {
	return q.size == 0
}

// Enqueue 入队
func (q *LinkedListQueue) Enqueue(data interface{}) {
	if q.tail == nil {
		q.tail = NewNode(data, nil)
		q.head = q.tail
		q.size++
	} else {
		q.tail.next = NewNode(data, nil)
		q.tail = q.tail.next
		q.size++
	}
}

// Dequeue 出队
func (q *LinkedListQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		panic("Queue is empty, no dequeue")
	}
	retNode := q.head
	q.head = q.head.next
	retNode.next = nil
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return retNode.data
}

// GetFront 查看队列元素
func (q *LinkedListQueue) GetFront() interface{} {
	if q.IsEmpty() {
		panic("Queue is empty, no get")
	}

	return q.head.data
}

func (q *LinkedListQueue) String() string {

	var buffer bytes.Buffer

	buffer.WriteString("Queue front [")

	current := q.head
	for current != nil {
		buffer.WriteString(fmt.Sprintf("%v -> ", current.data))
		current = current.next
	}

	buffer.WriteString("nil tail")

	return buffer.String()
}
