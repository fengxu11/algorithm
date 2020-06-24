package code

import "strconv"

// IsPalindrome 判断是否是回文数
func IsPalindrome(x int) bool {

	// 排除小于0的数
	if x < 0 {
		return false
	}

	// 转换为字符串
	xStr := strconv.Itoa(x)
	xStrReverse := make([]rune, 0)

	// 反转字符串
	for i := range xStr {
		xStrReverse = append(xStrReverse, rune(xStr[len(xStr)-1-i]))
	}

	// 进行比对
	for i := 0; i < len(xStr); i++ {
		if rune(xStr[i]) != xStrReverse[i] {
			return false
		}
	}

	return true
}
