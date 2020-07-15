package remove

import (
	"fmt"
	"testing"
)

func TestRemoveElementsByFor(t *testing.T) {

	testData1 := NewListNodeByArr([]int{1, 2, 3, 6, 4, 5, 6})
	fmt.Println("创建链表1: ", testData1)
	testData2 := NewListNodeByArr([]int{1})
	fmt.Println("创建链表2: ", testData2)

	t.Run("data 1", func(t *testing.T) {

		element := removeElementsByFor(testData1, 6)

		fmt.Println("删除后: ", element)
	})

	t.Run("data 2", func(t *testing.T) {

		element := removeElementsByFor(testData2, 1)

		fmt.Println("删除后: ", element)
	})
}

func TestRemoveElementsByDummyHead(t *testing.T) {

	testData1 := NewListNodeByArr([]int{1, 2, 3, 6, 4, 5, 6})
	fmt.Println("创建链表1: ", testData1)
	testData2 := NewListNodeByArr([]int{1})
	fmt.Println("创建链表2: ", testData2)

	t.Run("data 1", func(t *testing.T) {

		element := removeElementsByDummyHead(testData1, 6)

		fmt.Println("删除后: ", element)
	})

	t.Run("data 2", func(t *testing.T) {

		element := removeElementsByDummyHead(testData2, 1)

		fmt.Println("删除后: ", element)
	})
}
