package example

import (
	"fmt"
	"testing"
)

func TestLinkedListQueue(t *testing.T) {

	queue := NewLinkListQueue()

	for i := 0; i < 5; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)
	}

	fmt.Println("执行出队操作: ")
	queue.Dequeue()
	fmt.Println(queue)

	fmt.Println("查看队首元素: ", queue.GetFront())
}
