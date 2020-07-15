package example

import (
	"fmt"
	"testing"
	"time"
)

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

func TestStackTime(t *testing.T) {

	opCount := 1000000

	start := time.Now()

	stack := New()

	for i := 1; i <= opCount; i++ {
		stack.Push(i)
	}

	for i := 1; i <= opCount; i++ {
		stack.Pop()
	}

	t.Log(stack)

	end := time.Now()
	fmt.Println("Loop Queue time:", end.Sub(start).Seconds())
}
