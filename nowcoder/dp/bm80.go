package dp

/**
BM80 买卖股票的最好时机(一)
题目
题解(219)
讨论(332)
排行
面经new
简单  通过率：53.77%  时间限制：1秒  空间限制：256M
知识点
动态规划
贪心
描述
假设你有一个数组prices，长度为n，其中prices[i]是股票在第i天的价格，请根据这个价格数组，返回买卖股票能获得的最大收益
1.你可以买入一次股票和卖出一次股票，并非每天都可以买入或卖出一次，总共只能买入和卖出一次，且买入必须在卖出的前面的某一天
2.如果不能获取到任何利润，请返回0
3.假设买入卖出均无手续费

数据范围： 0 \le n \le 10^5 , 0 \le val \le 10^40≤n≤10
5
 ,0≤val≤10
4

要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
*/

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	min := 200000
	max := -1
	mPro := 0

	for _, v := range prices {
		if min > v {
			min = v
			max = v
		}
		if max < v {
			max = v
			if mPro < (max - min) {
				mPro = max - min
			}
		}
	}
	return mPro
}
