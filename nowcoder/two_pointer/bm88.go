package twopointer

/**
BM88 判断是否为回文字符串
描述
给定一个长度为 n 的字符串，请编写一个函数判断该字符串是否回文。如果是回文请返回true，否则返回false。

字符串回文指该字符串正序与其逆序逐字符一致。

数据范围：0 < n \le 10000000<n≤1000000
要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
*/

func judge(str string) bool {
	start := 0
	end := len(str) - 1

	for start <= end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}
	return true
}
