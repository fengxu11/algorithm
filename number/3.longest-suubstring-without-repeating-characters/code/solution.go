package code

import "math"

func lengthOfLongestSubstring(s string) int {

	if s == "" {
		return 0
	}

	start, end, res := 0, 0, 0

	lookup := map[byte]bool{}

	// 如果 start 和 end的值、都小于字符串的长度
	for start < len(s) && end < len(s) {

		// 从字符串中取出 当前end指向的
		// 然后从map中取这个字符串、
		// 如果没有 则记录下来、更新最大长度、同时 end指针往后走一步
		// 如果有、则从记录中删除之前出现过的字符、start指针往后走一步
		if ok := lookup[s[end]]; !ok {

			lookup[s[end]] = true
			// end - start 是当前子串的长度、然后+1 更新长度
			res = int(math.Max(float64(res), float64(end-start+1)))
			end++
		} else {

			delete(lookup, s[start])
			start++
		}

	}

	return res
}
