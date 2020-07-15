package recursion

import (
	"fmt"
	"testing"
)

func TestRemoveElementsByRecursion(t *testing.T) {

	testData1 := NewListNodeByArr([]int{1, 2, 3, 6, 4, 5, 6})
	fmt.Println("创建链表1: ", testData1)

	res := removeElements(testData1, 6)
	fmt.Println("链表val: ", res)
}
