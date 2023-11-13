package dp

/**

BM81 买卖股票的最好时机(二)
题目
题解(67)
讨论(75)
排行
面经new
中等  通过率：66.91%  时间限制：1秒  空间限制：256M
知识点
贪心
动态规划
描述
假设你有一个数组prices，长度为n，其中prices[i]是某只股票在第i天的价格，请根据这个价格数组，返回买卖股票能获得的最大收益
1. 你可以多次买卖该只股票，但是再次购买前必须卖出之前的股票
2. 如果不能获取收益，请返回0
3. 假设买入卖出均无手续费

数据范围： 1 \le n \le 1 \times 10^51≤n≤1×10
5
  ， 1 \le prices[i] \le 10^41≤prices[i]≤10
4

要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
进阶：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
*/

func maxProfit81(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	buy := prices[0]
	total := 0
	for k := 1; k < len(prices); k++ {
		if prices[k] < prices[k-1] {
			if prices[k-1]-buy > 0 {
				total += prices[k-1] - buy
			}
			buy = prices[k]
		}
	}
	if prices[len(prices)-1] > buy {
		total += prices[len(prices)-1] - buy
	}
	return total
}

func maxProfit812(prices []int) int {
	// write code here
	var res int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}
