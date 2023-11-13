package recursion

import (
	"sort"
	"strings"
)

/**
BM58 字符串的排列
题目
题解(251)
讨论(1k)
排行
面经new
中等  通过率：23.85%  时间限制：1秒  空间限制：256M
知识点
字符串
递归
描述
输入一个长度为 n 字符串，打印出该字符串中字符的所有排列，你可以以任意顺序返回这个字符串数组。
例如输入字符串ABC,则输出由字符A,B,C所能排列出来的所有字符串ABC,ACB,BAC,BCA,CBA和CAB。

数据范围：n < 10n<10
要求：空间复杂度 O(n!)O(n!)，时间复杂度 O(n!)O(n!)
输入描述：
输入一个字符串,长度不超过10,字符只包括大小写字母。
示例1
*/

func Permutation(str string) []string {
	if str == "" {
		return nil
	}
	strList := make([]string, 0, len(str))
	for _, v := range []byte(str) {
		strList = append(strList, string(v))
	}
	sort.Strings(strList)

	result := []string{}
	dfs(strList, 0, &result)

	return result
}

func dfs(strs []string, index int, result *[]string) {
	// 排列到最后一位，无需再处理，直接生成字符序列
	if index == len(strs)-1 {
		newStr := strings.Join(strs, "")
		*result = append(*result, newStr)
		return
	}

	for i := index; i < len(strs); i++ {
		// 剪枝
		if i != index && strs[i] == strs[index] {
			continue
		}
		// 递归处理下一层
		strs[i], strs[index] = strs[index], strs[i]
		dfs(strs, index+1, result)
	}
	for i := len(strs) - 1; i > index; i-- {
		strs[i], strs[index] = strs[index], strs[i]
	}
}
