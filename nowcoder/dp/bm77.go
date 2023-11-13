package dp

import (
	"container/list"
)

/*
*
BM77 最长的括号子串
较难  通过率：28.37%  时间限制：1秒  空间限制：256M
知识点
字符串
动态规划
栈
描述
给出一个长度为 n 的，仅包含字符 '(' 和 ')' 的字符串，计算最长的格式正确的括号子串的长度。

例1: 对于字符串 "(()" 来说，最长的格式正确的子串是 "()" ，长度为 2 .
例2：对于字符串 ")()())" , 来说, 最长的格式正确的子串是 "()()" ，长度为 4 .

字符串长度：0 \le n \le 5*10^50≤n≤5∗10
5
要求时间复杂度 O(n)O(n) ,空间复杂度 O(n)O(n).
*/
func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 0
	m := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i < 2 {
					dp[i] = 2
				} else {
					dp[i] = dp[i-2] + 2
				}
			}
			if dp[i-1] > 0 && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			if m < dp[i] {
				m = dp[i]
			}
		}
	}
	return m
}

func longestValidParentheses2(s string) int {
	if len(s) <= 1 {
		return 0
	}
	length, l, r, m := 0, 0, 0, 0
	for _, v := range []byte(s) {
		if v == '(' {
			l++
		} else {
			r++
			if l == r {
				length = l + r
				if m < length {
					m = length
				}
			} else if r > l {
				l, r = 0, 0
			}
		}
	}
	length, l, r = 0, 0, 0
	for k := len(s) - 1; k >= 0; k-- {
		v := s[k]
		if v == ')' {
			r++
		} else {
			l++
			if l == r {
				length = l + r
				if m < length {
					m = length
				}
			} else if r < l {
				l, r = 0, 0
			}
		}
	}
	return m
}

func longestValidParentheses3(s string) int {
	stack := list.New()
	stack.PushBack(-1)
	res := 0
	for k, v := range s {
		if v == '(' {
			stack.PushBack(k)
		} else {
			if stack.Len() > 1 {
				stack.Remove(stack.Back())
				res = max77(res, k-stack.Back().Value.(int))
			} else {
				stack.Remove(stack.Back())
				stack.PushBack(k)
			}
		}
	}
	return res
}
func max77(a ...int) int {
	m := a[0]
	for _, v := range a {
		if m < v {
			m = v
		}
	}
	return m
}
