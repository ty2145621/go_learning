package stack

import (
	"container/list"
)

/**
BM44 有效括号序列
题目
题解(240)
讨论(355)
排行
面经new
简单  通过率：35.75%  时间限制：1秒  空间限制：256M
知识点
栈
字符串
描述
给出一个仅包含字符'(',')','{','}','['和']',的字符串，判断给出的字符串是否是合法的括号序列
括号必须以正确的顺序关闭，"()"和"()[]{}"都是合法的括号序列，但"(]"和"([)]"不合法。

数据范围：字符串长度 0\le n \le 100000≤n≤10000
要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
*/

func isValid(str string) bool {
	s := list.New()
	for _, v := range []byte(str) {
		if v == '(' || v == '{' || v == '[' {
			s.PushBack(v)
		} else {
			if s.Len() == 0 {
				return false
			}
			top := s.Remove(s.Back()).(byte)
			if !isPair(top, v) {
				return false
			}
		}
	}
	if s.Len() > 0 {
		return false
	}
	return true
}

func isPair(a, b byte) bool {
	if a == '(' && b == ')' ||
		a == '[' && b == ']' ||
		a == '{' && b == '}' {
		return true
	}
	return false
}
