package recursion

import "testing"

func TestArraySum(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Log("sum array: ", ArraySum(arr))
}
