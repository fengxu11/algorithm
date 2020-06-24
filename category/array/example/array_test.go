package example

import (
	"testing"
)

func TestArray(t *testing.T) {

	arr := NewCapacity(10)
	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	t.Log(arr)

	arr.Add(1, 100)
	t.Log(arr)

	arr.AddFirst(-1)
	t.Log(arr)

	arr.Remove(2)
	t.Log(arr)

	arr.RemoveElement(4)
	t.Log(arr)

	arr.RemoveFirst()
	t.Log(arr)

	for i := 0; i < 4; i++ {
		arr.RemoveFirst()
		t.Log(arr)
	}
}
