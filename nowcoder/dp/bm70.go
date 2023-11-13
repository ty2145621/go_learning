package dp

import (
	"math"
	"sort"
)

/**
背包问题
*/

/**
BM70 兑换零钱(一)
题目
题解(79)
讨论(74)
排行
面经new
中等  通过率：43.19%  时间限制：2秒  空间限制：256M
知识点
动态规划
描述
给定数组arr，arr中所有的值都为正整数且不重复。每个值代表一种面值的货币，每种面值的货币可以使用任意张，再给定一个aim，代表要找的钱数，求组成aim的最少货币数。
如果无解，请返回-1.


数据范围：数组大小满足 0 \le n \le 100000≤n≤10000 ， 数组中每个数字都满足 0 < val \le 100000<val≤10000，0 \le aim \le 50000≤aim≤5000

要求：时间复杂度 O(n \times aim)O(n×aim) ，空间复杂度 O(aim)O(aim)。

*/

func minMoney(arr []int, aim int) int {
	// write code here
	// return minMoney1(arr, aim)
	return minMoney2(arr, aim)
}

func minMoney2(arr []int, aim int) int {
	if aim == 0 {
		return 0
	}
	dp := make([]int, aim+1)
	sort.Ints(arr)

	for i := 0; i <= aim; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= aim; i++ {
		for j := 0; j < len(arr); j++ {
			if i-arr[j] >= 0 {
				dp[i] = min70(dp[i], dp[i-arr[j]]+1)
			}
		}
	}
	if dp[aim] == math.MaxInt32 {
		return -1
	}
	return dp[aim]
}

func minMoney3(arr []int, aim int) int {
	// write code here
	// sort.Ints(arr)
	var dp = make([]int, aim+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt16
	}
	dp[0] = 0
	for i := 0; i < len(arr); i++ {
		for j := arr[i]; j <= aim; j++ {
			if dp[j-arr[i]]+1 < dp[j] {
				dp[j] = dp[j-arr[i]] + 1
			}
		}
	}
	if dp[aim] == math.MaxInt16 {
		return -1
	}
	return dp[aim]
}

func minMoney1(arr []int, aim int) int {
	if aim == 0 {
		return 0
	}
	dp := make([]int, aim+1)

	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < len(dp) {
			dp[v] = 1
		}
	}

	for k, _ := range dp {
		min := aim
		if dp[k] == 1 {
			continue
		}
		for _, v := range arr {
			if k-v > 0 && dp[k-v] != 0 {
				if min > dp[k-v] {
					min = dp[k-v]
				}
			}
		}
		if min != aim {
			dp[k] = min + 1
		}
	}
	if dp[aim] == 0 {
		return -1
	}
	return dp[aim]
}

func min70(a ...int) int {
	m := a[0]
	for _, v := range a {
		if m > v {
			m = v
		}
	}
	return m
}
