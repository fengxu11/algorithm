### No. 9
### Title: palindrome-number  回文数
### Tag: array
### Description: 
```

判断一个整数是否是 回文数。
回文数是指 从左往右 和 从右往左读都是一样的整数


```

### example
```

示例 1:
    输入: 121
    输出: true


示例 2:
    输入: -121
    输出: false
    解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。


示例 3:
    输入: 10
    输出: false
    解释: 从右向左读, 为 01 。因此它不是一个回文数。


进阶:
    你能不将整数转为字符串来解决这个问题吗？

```


### 解题思路

```

从题目和示例 可以得到以下几点 

    1. 回文数不可能是负数、如果是负数直接返回false
    2. 反转输入的数、与原来的数进行比对、值不同则直接返回false
    3. 如果不是负数、同时反转后的值 与 输入的值相等、则证明输入的值是一个回文数

```


### Code 
```go


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

	for i := 0; i < len(xStr); i++ {
		if rune(xStr[i]) != xStrReverse[i] {
			return false
		}
	}

	return true
}



```