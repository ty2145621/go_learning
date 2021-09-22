package code_set1

/**
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

输入：[1,8,6,2,5,4,8,3,7]
输出：49
*/

func checkMin11(height []int, l, r, max int) int {
	var min = height[l]
	if height[l] > height[r] {
		min = height[r]
	}
	s := min * (r - l)
	if max > s {
		return max
	}
	return s
}

func maxArea(height []int) int {
	var (
		l   = 0
		r   = len(height) - 1
		max = 0
	)

	for l < r {
		max = checkMin11(height, l, r, max)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}