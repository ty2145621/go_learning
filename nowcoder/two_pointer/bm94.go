package twopointer

/**
BM94 接雨水问题
题目
题解(151)
讨论(195)
排行
面经new
较难  通过率：42.24%  时间限制：1秒  空间限制：256M
知识点
双指针
单调栈
动态规划
描述
给定一个整形数组arr，已知其中所有的值都是非负的，将这个数组看作一个柱子高度图，计算按此排列的柱子，下雨之后能接多少雨水。(数组以外的区域高度视为0)
*/

func maxWater(arr []int) int64 {
	l, r := 0, len(arr)-1
	lmax, rmax := 0, 0
	t := 0
	for l < r {
		lmax = max(arr[l], lmax)
		rmax = max(arr[r], rmax)
		if lmax < rmax {
			t += lmax - arr[l]
			l++
		} else {
			t += rmax - arr[r]
			r--
		}
	}
	return int64(t)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
