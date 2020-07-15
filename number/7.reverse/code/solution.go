package code

// Integer 反转整数
func Integer(x int) int {

	// 如果是负数、取反、递归后再取反
	if x < 0 {
		return -Integer(-x)
	}

	var res int

	// 每次得到最后一位数字、并将其作为结果中的最高位
	for x != 0 {
		res = res*10 + x%10
		x /= 10
	}

	// 如果溢出、则返回0
	if res >= 0x7fffffff {
		return 0
	}

	return res
}
