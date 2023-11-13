package twopointer

import (
	"fmt"
)

/**
BM90 最小覆盖子串
描述

给出两个字符串 s 和 t，要求在 s 中找出最短的包含 t 中所有字符的连续子串。

数据范围：0 \le |S|,|T| \le100000≤∣S∣,∣T∣≤10000，保证s和t字符串中仅包含大小写英文字母
要求：进阶：空间复杂度 O(n)O(n) ， 时间复杂度 O(n)O(n)
例如：
S ="XDOYEZODEYXNZ"S="XDOYEZODEYXNZ"
T ="XYZ"T="XYZ"
找出的最短子串为"YXNZ""YXNZ".

注意：
如果 s 中没有包含 t 中所有字符的子串，返回空字符串 “”；
满足条件的子串可能有很多，但是题目保证满足条件的最短的子串唯一。

示例1
输入：
"XDOYEZODEYXNZ","XYZ"
复制
返回值：
"YXNZ"
*/

func minWindow(S string, T string) string {
	m := make(map[byte]int, 0)
	mCheck := make(map[byte]bool, 0)

	for _, v := range []byte(T) {
		m[v] -= 1
	}

	left := 0
	minLeft := -1
	minRight := len(S)
	k := 0
	for ; k < len(S); k++ {
		c := S[k]
		if _, ok := m[c]; !ok {
			continue
		}
		m[c]++
		if m[c] >= 0 {
			mCheck[c] = true
		}
		if len(mCheck) >= len(m) { // 找全了
			for left < k {
				cl := S[left]
				if mCount, ok := m[cl]; !ok {
					left++
					continue
				} else if mCount <= 0 {
					break
				}
				m[cl]--
				left++
			}
			minLeft, minRight = left, k
			break
		}
	}
	fmt.Println(len(mCheck), len(m), len(S), len(T), k, minLeft, minRight)

	for k = k + 1; k < len(S); k++ {
		c := S[k]
		if _, ok := m[c]; !ok {
			continue
		}
		m[c]++
		for left < k {
			cl := S[left]
			if mCount, ok := m[cl]; !ok {
				left++
				continue
			} else if mCount <= 0 {
				break
			}
			m[cl]--
			left++
		}
		if k-left < minRight-minLeft {
			minLeft, minRight = left, k
		}

	}
	if minLeft == -1 {
		return ""
	}
	return S[minLeft : minRight+1]
}
