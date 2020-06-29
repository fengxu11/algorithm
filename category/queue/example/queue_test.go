package example

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
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

	opCount := 100000

	t.Run("Array Queue", func(t *testing.T) {

		start := time.Now()

		queue := NewArrayQueue(10)

		for i := 0; i < opCount; i++ {
			queue.Enqueue(rand.Int())
		}

		for i := 0; i < opCount; i++ {
			queue.Dequeue()
		}

		end := time.Now()
		fmt.Println("Array Queue time", end.Sub(start).Seconds())
	})

	t.Run("Loop Queue", func(t *testing.T) {

		start := time.Now()

		queue := NewLoopQueue(10)

		for i := 0; i < opCount; i++ {
			queue.Enqueue(rand.Int())
		}

		for i := 0; i < opCount; i++ {
			queue.Dequeue()
		}

		end := time.Now()
		fmt.Println("Loop Queue time:", end.Sub(start).Seconds())
	})

}
