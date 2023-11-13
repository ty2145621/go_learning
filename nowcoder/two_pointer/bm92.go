package twopointer

/**
BM92 最长无重复子数组
题目
题解(285)
讨论(401)
排行
面经new
中等  通过率：32.76%  时间限制：1秒  空间限制：256M
知识点
哈希
双指针
数组
描述
给定一个长度为n的数组arr，返回arr的最长无重复元素子数组的长度，无重复指的是所有数字都不相同。
子数组是连续的，比如[1,3,5,7,9]的子数组有[1,3]，[3,5,7]等等，但是[1,3,7]不是子数组

数据范围：0\le arr.length \le 10^50≤arr.length≤10
5
 ，0 < arr[i] \le 10^50<arr[i]≤10
5
*/

func maxLength(arr []int) int {
	l, r := 0, 0
	max := 0
	m := make(map[int]int, 0)
	for r < len(arr) {
		m[arr[r]]++
		for m[arr[r]] > 1 && l < r {
			m[arr[l]]--
			l++
		}
		if max < r-l+1 {
			max = r - l + 1
		}
		r++
	}
	return max
}
