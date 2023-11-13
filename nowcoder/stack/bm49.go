package stack

/**
BM49 表达式求值
题目
题解(142)
讨论(149)
排行
面经new
中等  通过率：46.95%  时间限制：1秒  空间限制：256M
知识点
栈
递归
描述
请写一个整数计算器，支持加减乘三种运算和括号。

数据范围：0\le |s| \le 1000≤∣s∣≤100，保证计算结果始终在整型范围内

要求：空间复杂度： O(n)O(n)，时间复杂度 O(n)O(n)
示例1
*/

func solve(s string) int {
	// write code here
	stack := make([]int, 0)
	data := []byte(s)
	number := 0
	var operation byte = '+'
	result := 0
	for i := 0; i < len(data); i++ {
		// 0~9
		if data[i] >= '0' && data[i] <= '9' {
			number = 10*number + int(data[i]-'0')
		}

		// '('
		if data[i] == '(' {
			count := 1
			start, end := i+1, i+1
			for count != 0 {
				if data[end] == '(' {
					count++
				}
				if data[end] == ')' {
					count--
				}
				end++
			}
			i = end - 1
			number = solve(s[start : end-1])
		}

		// '+'/'-'/'*'
		if data[i] < '0' || data[i] > '9' || i == len(data)-1 {
			switch operation {
			case '+':
				stack = append(stack, number)
			case '-':
				stack = append(stack, -number)
			case '*':
				stack[len(stack)-1] = stack[len(stack)-1] * number
			}
			number = 0
			operation = data[i]
		}
	}

	for len(stack) != 0 {
		result += stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return result
}
