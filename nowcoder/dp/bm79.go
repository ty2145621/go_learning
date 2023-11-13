package dp

/**

BM79 打家劫舍(二)
题目
题解(36)
讨论(30)
排行
面经new
中等  通过率：53.12%  时间限制：2秒  空间限制：256M
知识点
动态规划
描述
你是一个经验丰富的小偷，准备偷沿湖的一排房间，每个房间都存有一定的现金，为了防止被发现，你不能偷相邻的两家，即，如果偷了第一家，就不能再偷第二家，如果偷了第二家，那么就不能偷第一家和第三家。沿湖的房间组成一个闭合的圆形，即第一个房间和最后一个房间视为相邻。
给定一个长度为n的整数数组nums，数组中的元素表示每个房间存有的现金数额，请你计算在不被发现的前提下最多的偷窃金额。

数据范围：数组长度满足 1 \le n \le 2\times10^5 \1≤n≤2×10
5
  ，数组中每个值满足 1 \le nums[i] \le 5000 \1≤nums[i]≤5000
*/

func rob79(nums []int) int {
	return rob791(nums)
}

func rob791(nums []int) int {
	if len(nums) <= 3 {
		return max79(nums...)
	}

	dp0 := make([]int, len(nums)+3) // 偷0
	dp1 := make([]int, len(nums)+3) // 不偷0
	dp0[0], dp0[1], dp0[2] = nums[0], nums[0], nums[0]
	dp1[0], dp1[1], dp1[2] = 0, nums[1], max79(nums[1], nums[2])
	for k := 2; k < len(nums)-1; k++ {
		dp0[k] = max79(dp0[k-2]+nums[k], dp0[k-1])
	}
	dp0[len(nums)-1] = dp0[len(nums)-2]
	for k := 2; k < len(nums); k++ {
		dp1[k] = max79(dp1[k-2]+nums[k], dp1[k-1])
	}
	return max79(dp0[len(nums)-1], dp1[len(nums)-1])
}

func max79(a ...int) int {
	m := a[0]
	for _, v := range a {
		if m < v {
			m = v
		}
	}
	return m
}
