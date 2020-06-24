package example

import "github.com/algorithm/category/array/example"

// IArrayQueue 数组队列动作接口
type IArrayQueue interface {
	GetSize() int
	IsEmpty() bool
	Enqueue(e interface{})
	Dequeue() interface{}
	GetFront() interface{}
}

// ArrayQueue 数组队列
type ArrayQueue struct {
	array *example.Array
}

// NewArrayQueue 创建一个数组队列
func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		array: example.New(),
	}
}

// GetSize 获取大小
func (q *ArrayQueue) GetSize() int {
	return q.array.GetSize()
}

// IsEmpty 判断是否为空
func (q *ArrayQueue) IsEmpty() bool {
	return q.array.IsEmpty()
}

// Enqueue 向队列尾端 添加元素
func (q *ArrayQueue) Enqueue(e interface{}) {
	q.array.AddLast(e)
}

// Dequeue 从队首获取元素
func (q *ArrayQueue) Dequeue() interface{} {
	return q.array.RemoveFirst()
}

// GetFront 获取第一个元素
func (q *ArrayQueue) GetFront() interface{} {
	return q.array.GetFirst()
}

// ILoopQueue 循环队列
type ILoopQueue interface {
}
