package dp

/**
BM71 最长上升子序列(一)
中等  通过率：41.40%  时间限制：1秒  空间限制：256M
知识点
动态规划
描述
给定一个长度为 n 的数组 arr，求它的最长严格上升子序列的长度。
所谓子序列，指一个数组删掉一些数（也可以不删）之后，形成的新数组。例如 [1,5,3,7,3] 数组，其子序列有：[1,3,3]、[7] 等。但 [1,6]、[1,3,5] 则不是它的子序列。
我们定义一个序列是 严格上升 的，当且仅当该序列不存在两个下标 ii 和 jj 满足 i<ji<j 且 arr_i \geq arr_jarr
i
​
 ≥arr
j
​
 。
数据范围： 0\leq n \leq 10000≤n≤1000
要求：时间复杂度 O(n^2)O(n
2
 )， 空间复杂度 O(n)O(n)
示例1
输入：
[6,3,1,5,2,3,7]
返回值：
4
说明：
该数组最长上升子序列为 [1,2,3,7] ，长度为4
*/

func LIS(arr []int) int {
	return LIS1(arr)
}

func LIS1(arr []int) int {
	dp := make([]int, len(arr)+1)
	for k, v := range arr {
		max := 0
		for i := 0; i < k; i++ {
			if v > arr[i] && dp[i] > max {
				max = dp[i]
			}
		}
		dp[k] = max + 1
	}
	m := 0
	for _, v := range dp {
		if v > m {
			m = v
		}
	}
	return m
}
