package code

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
