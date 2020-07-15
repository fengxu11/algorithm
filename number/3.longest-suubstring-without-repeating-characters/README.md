### No. 3
### Title: longest-substring-without-repeating-characters  无重复字符的最长子串
### Tag: array
### Description: 
```

给定一个字符串、请你找出其中不含有重复字符的 “最长子串” 的长度


```

### example
```
示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符额最长子串是 “abc”, 所以它的长度为3


示例 2:

输入: "bbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b", 所以它的长度是1 


示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 “wke”, 所以它的长度是 3


注意:
    你的答案必须是 ”子串“ 的长度, "pwke" 是一个子序列, 不是一个子串
```


### 解决方案

```

题目理解: 
    1. 子串必须 要连续的、 子序列可以不连续、比如示例3
    2. 题目要求的是 不包括重复字符串的子串
    3. 题目要求的是 满足条件的子串中 最长的那个子串的长度



解题1: 
    利用窗口滑动的思想来解决这个问题 (sliding window)
    1. 每次遍历一个字符、记录字符、没记录一次就把最大长度更新一次
    2. 每次拿到一个字符就去 记录里面找、如果有把之前的记录删掉
    3. 这样 遍历到最后一个、条件结束、最大长度就是结果值

start
end
  |
  |
  v

  a  b  c  a  b  c  b  b

    end指针一直不停地往前走
    只要当前子串 s[start:end+1] 不满足 无重复字符条件的时候
    我们就让 start指针往前走、直到满足条件为止、每次满足条件后更新一下最大长度、最后就是结果值


```


### Code

```go


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





```