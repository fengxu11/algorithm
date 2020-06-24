package example

import (
	"bytes"
	"fmt"

	"github.com/algorithm/category/common"
)

const failMessage string = "fail, index < 0 or index > arr size"

// Array Encapsulating an array with better performance based on slice
type Array struct {
	data []interface{}
	size int
}

// NewCapacity create an array based on capacity
func NewCapacity(capacity int) *Array {
	return &Array{
		data: make([]interface{}, capacity),
	}
}

// New create an array
func New() *Array {
	return &Array{
		data: make([]interface{}, 0),
	}
}

// GetCap get array capacity
func (array *Array) GetCap() int {
	return len(array.data)
}

// GetSize get array size
func (array *Array) GetSize() int {
	return array.size
}

// IsEmpty judge whether the array is empty
func (array *Array) IsEmpty() bool {
	return array.size == 0
}

// Add insert e into index
func (array *Array) Add(index int, e interface{}) {
	if index < 0 || index > array.size {
		panic(failMessage)
	}

	if array.size == len(array.data) {
		cap := array.size
		if cap == 0 {
			cap = 1
		}
		array.resize(2 * cap)
	}

	for i := array.size - 1; i >= index; i-- {
		array.data[i+1] = array.data[i]
	}

	array.data[index] = e
	array.size++
}

// AddLast add last
func (array *Array) AddLast(e interface{}) {
	array.Add(array.size, e)
}

// AddFirst add first
func (array *Array) AddFirst(e interface{}) {
	array.Add(0, e)
}

// Get get the element of index location
func (array *Array) Get(index int) interface{} {
	if index < 0 || index >= array.size {
		panic(failMessage)
	}
	return array.data[index]
}

// GetLast 获取最后一个元素
func (array *Array) GetLast() interface{} {
	return array.Get(array.size - 1)
}

// GetFirst 获取第一个元素
func (array *Array) GetFirst() interface{} {
	return array.Get(0)
}

// Set update the element of index location
func (array *Array) Set(index int, e interface{}) {
	if index < 0 || index >= array.size {
		panic(failMessage)
	}
	array.data[index] = e
}

// Contains find if element E exists
func (array *Array) Contains(e interface{}) bool {
	for i := 0; i < array.size; i++ {
		if common.Compare(array.data[i], e) == 0 {
			return true
		}
	}
	return false
}

// Find index for element E
// -1 not exists
func (array *Array) Find(e interface{}) int {
	for i := 0; i < array.size; i++ {
		if common.Compare(array.data[i], e) == 0 {
			return i
		}
	}
	return -1
}

// Remove the element at index position from the array
// return removed element
func (array *Array) Remove(index int) interface{} {
	if index < 0 || index >= array.size {
		panic(failMessage)
	}

	e := array.data[index]
	for i := index + 1; i < array.size; i++ {
		array.data[i-1] = array.data[i]
	}
	array.size--
	array.data[array.size] = nil

	// 考虑边界条件，避免长度为 1 时，resize 为 0
	// 右移1位 等同于除以2、同等右移 两位等同于除以4
	if array.size == len(array.data)>>2 && len(array.data)>>1 != 0 {
		array.resize(len(array.data) >> 1)
	}
	return e
}

// RemoveFirst Delete first element
// return removed element
func (array *Array) RemoveFirst() interface{} {
	return array.Remove(0)
}

// RemoveLast Delete last element
// return removed element
func (array *Array) RemoveLast() interface{} {
	return array.Remove(array.size - 1)
}

// RemoveElement Remove an element E from the array
// return bool
func (array *Array) RemoveElement(e interface{}) bool {
	index := array.Find(e)
	if index == -1 {
		return false
	}

	array.Remove(index)
	return true
}

// RemoveAllElement Remove all elements e from the array
// return bool
func (array *Array) RemoveAllElement(e interface{}) bool {
	if array.Find(e) == -1 {
		return false
	}

	for i := 0; i < array.size; i++ {
		if common.Compare(array.data[i], e) == 0 {
			array.Remove(i)
		}
	}
	return true
}

// Array capacity expansion
func (array *Array) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity)
	for i := 0; i < array.size; i++ {
		newData[i] = array.data[i]
	}

	array.data = newData
}

// Overriding the string method of array
func (array *Array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", array.size, len(array.data)))
	buffer.WriteString("[")
	for i := 0; i < array.size; i++ {
		buffer.WriteString(fmt.Sprint(array.data[i]))
		if i != (array.size - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
