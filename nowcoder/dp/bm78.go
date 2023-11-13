package dp

/**
BM78 打家劫舍(一)
中等  通过率：49.57%  时间限制：1秒  空间限制：256M
知识点
动态规划
描述
你是一个经验丰富的小偷，准备偷沿街的一排房间，每个房间都存有一定的现金，为了防止被发现，你不能偷相邻的两家，即，如果偷了第一家，就不能再偷第二家；如果偷了第二家，那么就不能偷第一家和第三家。
给定一个整数数组nums，数组中的元素表示每个房间存有的现金数额，请你计算在不被发现的前提下最多的偷窃金额。

数据范围：数组长度满足 1 \le n \le 2\times 10^5\1≤n≤2×10
5
   ，数组中每个值满足 1 \le num[i] \le 5000 \1≤num[i]≤5000
*/

func rob(nums []int) int {
	return rob1(nums)
}

func rob1(nums []int) int {

	dp := make([]int, len(nums)+1)
	dp[0] = nums[0]
	for k := 1; k < len(nums); k++ {
		dp[k+1] = max78(dp[k-1]+nums[k], dp[k])
	}
	return dp[len(nums)]
}

func max78(a, b int) int {
	if a > b {
		return a
	}
	return b
}
