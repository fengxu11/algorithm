package example

import (
	"fmt"
	"testing"
)

func TestArrayQueue(t *testing.T) {

	// 1. 创建队列
	queue := NewArrayQueue(10)
	t.Log("队列长度: ", queue.GetSize())

	// 2. 往队列中添加
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	t.Log("队列长度: ", queue.GetSize())

	// 3. 取出队首元素
	t.Log("取出队首元素: ", queue.Dequeue())

	// 4. 查看队首元素
	t.Log("查看队守元素: ", queue.GetFront())
	t.Log("队列长度: ", queue.GetSize())

}

func TestLoopQueue(t *testing.T) {

	queue := NewLoopQueue(10)

	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)
		if i%3 == 2 {
			queue.Dequeue()
			fmt.Println(queue)
		}
	}

}

//  测试数组队列和循环队列的速度
func TestCompareOfArrayQueueAndLoopQueue(t *testing.T) {

	
}
