package twopointer

/**
BM93 盛水最多的容器
双指针
数组
描述
给定一个数组height，长度为n，每个数代表坐标轴中的一个点的高度，height[i]是在第i点的高度，请问，从中选2个高度与x轴组成的容器最多能容纳多少水
1.你不能倾斜容器
2.当n小于2时，视为不能形成容器，请返回0
3.数据保证能容纳最多的水不会超过整形范围，即不会超过231-1

数据范围:
0<=height.length<=10^50<=height.length<=10
5

0<=height[i]<=10^40<=height[i]<=10
4

如输入的height为[1,7,3,2,4,5,8,2,7]，那么如下图:
*/

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for left < right {
		c := min(height[left], height[right]) * (right - left)
		if max < c {
			max = c
		}
		if height[left] < height[right] {
			left++
		} else if height[left] >= height[right] {
			right--
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
