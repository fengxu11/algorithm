package example

import (
	"bytes"
	"fmt"

	"github.com/algorithm/category/array/example"
)

// S 栈的接口
type S interface {
	GetSize() int
	IsEmpty() bool
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
}

// Stack 栈的结构体
type Stack struct {
	array *example.Array
}

// GetSize 获取栈的大小
func (s *Stack) GetSize() int {
	return s.array.GetSize()
}

// IsEmpty 是否为空
func (s *Stack) IsEmpty() bool {
	return s.array.IsEmpty()
}

// Push 添加一个元素
func (s *Stack) Push(e interface{}) {
	s.array.AddLast(e)
}

// Pop 取出栈顶的元素
func (s *Stack) Pop() interface{} {
	return s.array.RemoveLast()
}

// Peek 查看栈顶的元素
func (s *Stack) Peek() interface{} {
	return s.array.GetLast()
}

func (s *Stack) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Stack: ")
	buffer.WriteString("[")
	for i := 0; i < s.array.GetSize(); i++ {
		buffer.WriteString(fmt.Sprint(s.array.Get(i)))
		// 加入逗号作为分隔
		if s.array.GetSize()-1 != i {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] stack top")
	return buffer.String()
}

// New 创建
func New() S {
	return &Stack{
		array: example.New(),
	}
}
