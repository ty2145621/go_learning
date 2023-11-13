package dp

/**
BM64 最小花费爬楼梯
知识点
动态规划
描述
给定一个整数数组 cost \cost  ，其中 cost[i]\cost[i]  是从楼梯第i \i 个台阶向上爬需要支付的费用，下标从0开始。一旦你支付此费用，即可选择向上爬一个或者两个台阶。

你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。

请你计算并返回达到楼梯顶部的最低花费。

数据范围：数组长度满足 1 \le n \le 10^5 \1≤n≤10
5
   ，数组中的值满足 1 \le cost_i \le 10^4 \1≤cost
i
​
 ≤10
4

*/

func minCostClimbingStairs(cost []int) int {
	// write code here
	return minCostClimbingStairs1(cost)
}

func minCostClimbingStairs1(cost []int) int {
	if len(cost) <= 2 {
		return getMin64(cost...)
	}
	a, b, c := cost[0], cost[1], getMin64(cost[0], cost[1])
	for i := 2; i < len(cost); i++ {
		c, a, b = getMin64(a+cost[i], b+cost[i]), b, getMin64(a+cost[i], b+cost[i])
	}
	return getMin64(c, a)
}

func getMin64(a ...int) int {
	m := a[0]
	for _, v := range a {
		if m > v {
			m = v
		}
	}
	return m
}
