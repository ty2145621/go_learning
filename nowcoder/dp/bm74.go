package dp

import (
	"strconv"
	"strings"
)

/**
BM74 数字字符串转化成IP地址
题目
题解(60)
讨论(151)
排行
面经new
中等  通过率：38.71%  时间限制：1秒  空间限制：256M
知识点
字符串
回溯
描述
现在有一个只包含数字的字符串，将该字符串转化成IP地址的形式，返回所有可能的情况。
例如：
给出的字符串为"25525522135",
返回["255.255.22.135", "255.255.221.35"]. (顺序没有关系)

数据范围：字符串长度 0 \le n \le 120≤n≤12
要求：空间复杂度 O(n!)O(n!),时间复杂度 O(n!)O(n!)

注意：ip地址是由四段数字组成的数字序列，格式如 "x.x.x.x"，其中 x 的范围应当是 [0,255]。
*/

func restoreIpAddresses(s string) []string {
	// write code here
	result := make([]string, 0)
	result = restoreIpAddresses1(s, 4, []string{}, result)
	return result
}

func isValid(s string) bool {
	if len(s) == 1 {
		return true
	}
	if s[0] == '0' {
		return false
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if i < 0 || i > 255 {
		return false
	}
	return true
}

func restoreIpAddresses1(s string, n int, arr []string, result []string) []string {
	if n == 0 {
		if len(s) == 0 {
			result = append(result, strings.Join(arr, "."))
		}
		return result
	}
	for i := 1; i <= 3 && i <= len(s); i++ {
		if isValid(s[0:i]) {
			result = restoreIpAddresses1(s[i:], n-1, append(arr, s[0:i]), result)
		}
	}
	return result
}
