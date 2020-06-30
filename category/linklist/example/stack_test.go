package example

import (
	"fmt"
	"testing"
	"time"
)

func TestLinkedListStack(t *testing.T) {

	opCount := 1000000

	start := time.Now()

	stack := NewLinkedListStack()
	t.Log(stack)

	for i := 0; i < opCount; i++ {
		stack.Push(i)
	}

	for i := 0; i < opCount; i++ {
		stack.Pop()
	}

	fmt.Println(stack)

	end := time.Now()
	fmt.Println("Loop Queue time:", end.Sub(start).Seconds())
}
