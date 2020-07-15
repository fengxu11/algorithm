package recursion

// ArraySum 数组求和
func ArraySum(arr []int) int {
	return sum(arr, 0)
}

func sum(arr []int, l int) int {

	// 递归结束条件
	if l == len(arr) {
		return 0
	}
	
	return arr[l] + sum(arr, l + 1)
}