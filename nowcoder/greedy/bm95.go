// Package greedy 贪心算法
package greedy

/**
BM95 分糖果问题
题目
题解(53)
讨论(63)
排行
面经new
较难  通过率：30.92%  时间限制：2秒  空间限制：256M
知识点
贪心
描述
一群孩子做游戏，现在请你根据游戏得分来发糖果，要求如下：

1. 每个孩子不管得分多少，起码分到一个糖果。
2. 任意两个相邻的孩子之间，得分较多的孩子必须拿多一些糖果。(若相同则无此限制)

给定一个数组 arrarr 代表得分数组，请返回最少需要多少糖果。

要求: 时间复杂度为 O(n)O(n) 空间复杂度为 O(n)O(n)

数据范围： 1 \le n \le 1000001≤n≤100000 ，1 \le a_i \le 10001≤a
i
​
 ≤1000
*/

func candy(arr []int) int {
	res := make([]int, len(arr))
	for k := 0; k < len(arr); k++ {
		res[k] = 1
	}
	for k := 1; k < len(arr); k++ {
		if arr[k] > arr[k-1] {
			res[k] = res[k-1] + 1
		}
	}
	for k := len(arr) - 2; k >= 0; k-- {
		if arr[k] > arr[k+1] {
			res[k] = max(res[k], res[k+1]+1)
		}
	}
	t := 0
	for k := 0; k < len(arr); k++ {
		t += res[k]
	}
	return t
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
