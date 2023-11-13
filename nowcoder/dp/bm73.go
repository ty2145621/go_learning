package dp

/**
BM73 最长回文子串
题目
题解(252)
讨论(402)
排行
面经new
中等  通过率：40.26%  时间限制：1秒  空间限制：256M
知识点
字符串
动态规划
描述
对于长度为n的一个字符串A（仅包含数字，大小写英文字母），请设计一个高效算法，计算其中最长回文子串的长度。


数据范围： 1 \le n \le 10001≤n≤1000
要求：空间复杂度 O(1)O(1)，时间复杂度 O(n^2)O(n
2
 )
进阶:  空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
*/

func getLongestPalindrome(str string) int {
	n := len(str)
	if n < 2 {
		return n
	}
	maxLen := 1

	// 奇数回文
	for i := 0; i < n; i++ {
		for j, k := i-1, i+1; j >= 0 && k < n && str[j] == str[k]; j, k = j-1, k+1 {
			if maxLen < k-j+1 {
				maxLen = k - j + 1
			}
		}
	}

	// 偶数回文
	for i := 0; i < n; i++ {
		for j, k := i, i+1; j >= 0 && k < n && str[j] == str[k]; j, k = j-1, k+1 {
			if maxLen < k-j+1 {
				maxLen = k - j + 1
			}
		}
	}

	return maxLen
}
