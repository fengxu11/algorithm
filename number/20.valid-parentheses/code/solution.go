package code

func isValid(s string) bool {

	stack := make([]rune, 0)

	backendMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, v := range s {

		// 如果是左括号、执行入栈操作
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		}

		// 如果是右括号、开始进行判断
		if v == ')' || v == '}' || v == ']' {

			// 栈为空 则字符串无效
			if len(stack) == 0 {
				return false
			}

			// 取出栈顶元素 与 右括号的对应的左括号取出来 进行匹配
			// 匹配成功则元素出栈、失败则字符串无效
			if stack[len(stack)-1] == backendMap[v] {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		}

	}

	return len(stack) == 0
}
