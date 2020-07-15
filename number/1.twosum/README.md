
### No. 1
### Title: two-sum 
### Tag: array
### Description: 
```

给定一个整数数组 numbers 和 target, 请你在该数组中找出: 
" 和为 " target的那两个整数, 并返回它们的数组下标

你可以假设每种输入只会对应一个答案、但是数组中同一个元素 不能使用两遍

```

### example
```
给定 numbers = [2, 7, 11, 15], target = 9

因为 numbers[0] + numbers[1] = 2 + 7 = 9
所以返回 [0, 1]

```


### 暴力解题法

```
双重遍历每一个元素、 拿到 i 和 j位置上的元素 与 target进行匹配、如果一直则返回

时间复杂度是 O(n2)
```

### Code
```go

// ByForce 暴力解法 - 双循环
func ByForce(numbers []int, target int) []int {

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

```


### 哈希表解法

```

1. 遍历数组
2. 判断 map中是否存在 target - 当前遍历到的值  m[target-number]
3. 如果有则返回 当前遍历到的值的下标 和 map中 值对应的下标
3. 如果没有则 把值作为key、下标作为value 放入map中

```

### Code
```go

// ByMap 哈希表解题法
func ByMap(numbers []int, target int) []int {

	m := make(map[int]int, 0)
	for key, number := range numbers {

		if val, ok := m[target-number]; ok {
			return []int{val, key}
		}
		m[number] = key
	}

	return nil
}

```