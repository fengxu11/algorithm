package example

import "testing"

func TestStack(t *testing.T) {
	stack := New()

	t.Log(stack.IsEmpty())

	t.Log(stack)

	for i := 1; i <= 10; i++ {
		stack.Push(i)
	}
	t.Log(stack)

	stack.Pop()

	t.Log(stack)

	t.Log(stack.Peek())

	t.Log(stack.GetSize())

	t.Log(stack.IsEmpty())
}
