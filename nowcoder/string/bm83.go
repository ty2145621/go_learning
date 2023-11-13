package string

import (
	"strings"
)

/**
BM83 字符串变形
题目
题解(90)
讨论(190)
排行
面经new
简单  通过率：29.15%  时间限制：1秒  空间限制：256M
知识点
字符串
描述
对于一个长度为 n 字符串，我们需要对它做一些变形。

首先这个字符串中包含着一些空格，就像"Hello World"一样，然后我们要做的是把这个字符串中由空格隔开的单词反序，同时反转每个字符的大小写。

比如"Hello World"变形后就变成了"wORLD hELLO"。

数据范围: 1\le n \le 10^61≤n≤10
6
  , 字符串中包括大写英文字母、小写英文字母、空格。
进阶：空间复杂度 O(n)O(n) ， 时间复杂度 O(n)O(n)
输入描述：
给定一个字符串s以及它的长度n(1 ≤ n ≤ 10^6)
返回值描述：
请返回变形后的字符串。题目保证给定的字符串均由大小写字母和空格构成。
*/

func trans(s string, n int) string {
	arr := strings.Split(s, " ")
	res := make([]string, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		res = append(res, updown(arr[i]))
	}
	return strings.Join(res, " ")
}

func updown(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return r + 'A' - 'a'
		} else {
			return r - 'A' + 'a'
		}
	}, s)
}
