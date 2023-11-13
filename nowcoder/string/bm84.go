package string

/*
*
BM84 最长公共前缀
题目
题解(116)
讨论(201)
排行
面经new
简单  通过率：37.03%  时间限制：1秒  空间限制：256M
知识点
字符串
描述
给你一个大小为 n 的字符串数组 strs ，其中包含n个字符串 , 编写一个函数来查找字符串数组中的最长公共前缀，返回这个公共前缀。

数据范围： 0 \le n \le 50000≤n≤5000， 0 \le len(strs_i) \le 50000≤len(strs
i
​

	)≤5000

进阶：空间复杂度 O(1)O(1)，时间复杂度 O(n*len)O(n∗len)
*/
func longestCommonPrefix2(strs []string) string {
	if strs == nil {
		return ""
	}
	strsLen := len(strs)
	if strsLen == 1 {
		return strs[0]
	}

	preLeft := longestCommonPrefix(strs[:strsLen/2])
	preRight := longestCommonPrefix(strs[strsLen/2:])

	common := -1
	for i, j := 0, 0; i < len(preLeft) && j < len(preRight) && preLeft[i] == preRight[i]; i, j, common = i+1, j+1, common+1 {
	}
	return preLeft[:common+1]
}

func longestCommonPrefix(strs []string) string {
	l := getMinLen(strs)
	if l == 0 {
		return ""
	}
	pos := 0
	for ; pos < l; pos++ {
		b := strs[0][pos]
		for i := 1; i < len(strs); i++ {
			if strs[i][pos] != b {
				return strs[0][:pos]
			}
		}
	}
	return strs[0][:pos]
}

func getMinLen(strs []string) int {
	if len(strs) == 0 {
		return 0
	}
	l := len(strs[0])
	for i := 0; i < len(strs); i++ {
		if l > len(strs[i]) {
			l = len(strs[i])
		}
	}
	return l
}
