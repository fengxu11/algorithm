package example

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/algorithm/category/array/example"
)

// IQueue 数组队列动作接口
type IQueue interface {
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

// LoopQueue 循环队列
type LoopQueue struct {
	Data  []interface{}
	Front int
	Tail  int
	Size  int
}

// NewLoopQueue 创建循环队列
func NewLoopQueue(capacity int) *LoopQueue {
	return &LoopQueue{
		Data:  make([]interface{}, capacity+1),
		Front: 0,
		Tail:  0,
		Size:  0,
	}
}

// GetCapacity 获取队列长度
func (q *LoopQueue) GetCapacity() int {
	return len(q.Data) - 1
}

// IsEmpty 判断循环队列是否为空
func (q *LoopQueue) IsEmpty() bool {
	return q.Front == q.Tail
}

// GetSize 获取循环队列大小
func (q *LoopQueue) GetSize() int {
	return q.Size
}

// Enqueue 入队
func (q *LoopQueue) Enqueue(e interface{}) {

	// 当tail + 1 余上 数组长度 等于 front、说明队列没有空间了、此时需要扩容
	if ((q.Tail + 1) % len(q.Data)) == q.Front {
		q.resize(q.GetCapacity() * 2)
	}
	q.Data[q.Tail] = e
	q.Tail = (q.Tail + 1) % len(q.Data)
	q.Size++
}

func (q *LoopQueue) resize(newCapacity int) {

	newData := make([]interface{}, newCapacity+1)
	for i := 0; i < q.Size; i++ {
		// 原有的data 和 新的data 不一样
		// 所以就导致了、 front的位置不一样、所以这里要做一个索引的偏移
		newData[i] = q.Data[(i+q.Front)%len(q.Data)]
	}

	q.Data = newData
	q.Front = 0
	q.Tail = q.Size
}

// Dequeue 出队
func (q *LoopQueue) Dequeue() interface{} {

	// 如果队列为空 则panic
	if q.IsEmpty() {
		panic(errors.New("Cannot dequeue from an empty queue"))
	}

	// 拿到出队的元素
	res := q.Data[q.Front]
	// front指向的数据 置空
	q.Data[q.Front] = nil
	// 维护front
	q.Front = (q.Front + 1) % len(q.Data)
	// 长度减1
	q.Size--

	// 数组缩容
	if q.Size == q.GetCapacity()/4 && q.GetCapacity()/2 != 0 {
		q.resize(q.GetCapacity() / 2)
	}

	return res
}

// GetFront 查看队首元素
func (q *LoopQueue) GetFront() interface{} {
	if q.IsEmpty() {
		panic("Queue is empty")
	}
	return q.Data[q.Front]
}

// Overriding the string method of array
func (q *LoopQueue) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Queue: size = %d, capacity = %d\n", q.Size, q.GetCapacity()))
	buffer.WriteString("front [")
	for i := q.Front; i != q.Tail; i = (i + 1) % len(q.Data) {
		buffer.WriteString(fmt.Sprint(q.Data[i]))
		if (i+1)%len(q.Data) != q.Tail {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")

	return buffer.String()
}
