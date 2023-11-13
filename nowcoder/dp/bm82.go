package dp

/**
BM82 买卖股票的最好时机(三)
题目
题解(44)
讨论(55)
排行
面经new
较难  通过率：49.71%  时间限制：1秒  空间限制：256M
知识点
动态规划
描述
假设你有一个数组prices，长度为n，其中prices[i]是某只股票在第i天的价格，请根据这个价格数组，返回买卖股票能获得的最大收益
1. 你最多可以对该股票有两笔交易操作，一笔交易代表着一次买入与一次卖出，但是再次购买前必须卖出之前的股票
2. 如果不能获取收益，请返回0
3. 假设买入卖出均无手续费

数据范围：1 \le n \le 10^51≤n≤10
5
 ，股票的价格满足 1 \le val\le 10^41≤val≤10
4

要求: 空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
进阶：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
*/

func maxProfit82(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp := make([][]int, 0)
	for i := 0; i < 5; i++ {
		dp = append(dp, make([]int, len(prices)+1))
	}
	dp[0][0] = 0
	dp[1][0] = -prices[0]
	for k := 1; k < len(prices); k++ {
		dp[0][k] = dp[0][k-1]
		dp[1][k] = max82(dp[1][k-1], dp[0][k-1]-prices[k])
		dp[2][k] = max82(dp[2][k-1], dp[1][k-1]+prices[k])
		dp[3][k] = max82(dp[3][k-1], dp[2][k-1]-prices[k])
		dp[4][k] = max82(dp[4][k-1], dp[3][k-1]+prices[k])
	}
	return max82(dp[4][len(prices)-1], dp[2][len(prices)-1])
}

func max82(a ...int) int {
	m := a[0]
	for _, v := range a {
		if m < v {
			m = v
		}
	}
	return m
}

func maxProfit822(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, prices[i]+buy1)
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, prices[i]+buy2)
	}
	return max(sell1, sell2)
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
