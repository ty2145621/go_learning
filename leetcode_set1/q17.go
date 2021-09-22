package leetcode_set1

import "fmt"

/**
17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例:
输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
*/

func letterCombinations(digits string) []string {
	m := map[uint8]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	if len(digits) == 0 {
		return make([]string, 0)
	}

	total := 1
	for _, v := range []byte(digits) {
		if v == '7' || v == '9' {
			total *= 4
		} else {
			total *= 3
		}
	}

	l := len(digits)
	ret := make([][]uint8, total)
	for k, _ := range ret {
		ret[k] = make([]uint8, l)
	}

	tl := 1
	for k1, v := range []byte(digits) {
		fmt.Printf("k1:%d\n", k1)
		t := 1
		if v == '7' || v == '9' {
			t = 4
		} else {
			t = 3
		}
		for k2, _ := range ret {
			ret[k2][k1] = m[v][(k2/tl)%t]
		}
		tl *= t
	}

	retStrArr := make([]string, 0)
	for _, v := range ret {
		retStrArr = append(retStrArr, string(v))
	}
	return retStrArr
}
