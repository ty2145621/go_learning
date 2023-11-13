package dp

/**
BM72 连续子数组的最大和
题目
题解(244)
讨论(1k)
排行
面经new
简单  通过率：40.49%  时间限制：1秒  空间限制：256M
知识点
动态规划
贪心
描述
输入一个长度为n的整型数组array，数组中的一个或连续多个整数组成一个子数组，子数组最小长度为1。求所有子数组的和的最大值。
数据范围:
1 <= n <= 2\times10^51<=n<=2×10
5

-100 <= a[i] <= 100−100<=a[i]<=100

要求:时间复杂度为 O(n)O(n)，空间复杂度为 O(n)O(n)
进阶:时间复杂度为 O(n)O(n)，空间复杂度为 O(1)O(1)
*/

func FindGreatestSumOfSubArray(array []int) int {
	return FindGreatestSumOfSubArray1(array)
}

func FindGreatestSumOfSubArray1(array []int) int {
	m := 0
	max := -101
	for _, v := range array {
		m += v
		if max < m {
			max = m
		}
		if m < 0 {
			m = 0
		}

	}
	return max
}
