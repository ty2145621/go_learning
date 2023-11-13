package dp

/**
BM63 跳台阶
题目
题解(385)
讨论(2k)
排行
面经new
简单  通过率：40.97%  时间限制：1秒  空间限制：256M
知识点
递归
动态规划
记忆化搜索
描述
一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个 n 级的台阶总共有多少种跳法（先后次序不同算不同的结果）。

数据范围：1 \leq n \leq 401≤n≤40
要求：时间复杂度：O(n)O(n) ，空间复杂度： O(1)O(1)
*/

func jumpFloor(number int) int {
	// write code here
	return jumpFloor1(number)
}

func jumpFloor1(n int) int {
	if n < 3 {
		return n
	}
	a, b, c := 1, 2, 3
	for i := 3; i <= n; i++ {
		c, a, b = a+b, b, a+b
	}
	return c
}
