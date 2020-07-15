package code

import (
	"fmt"
	"testing"
)

func TestByForce(t *testing.T) {

	numbers := []int{2, 7, 11, 15}
	target := 26

	// 暴力 解题法
	fmt.Println("by force result: ", ByForce(numbers, target))

	// map 解题法
	fmt.Println("by map result: ", ByMap(numbers, target))

}
