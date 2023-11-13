package dp

/**
BM68 矩阵的最小路径和
题目
题解(105)
讨论(155)
排行
面经new
中等  通过率：52.90%  时间限制：1秒  空间限制：256M
知识点
数组
动态规划
描述
给定一个 n * m 的矩阵 a，从左上角开始每次只能向右或者向下走，最后到达右下角的位置，路径上所有的数字累加起来就是路径和，输出所有的路径中最小的路径和。

数据范围: 1 \le n,m\le 5001≤n,m≤500，矩阵中任意值都满足 0 \le a_{i,j} \le 1000≤a
i,j
​
 ≤100
要求：时间复杂度 O(nm)O(nm)

例如：当输入[[1,3,5,9],[8,1,3,4],[5,0,6,1],[8,8,4,0]]时，对应的返回值为12，
所选择的最小累加和路径如下图所示：
*/

func minPathSum(matrix [][]int) int {
	return minPathSum1(matrix)
}

func minPathSum1(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	dp := make([][]int, 0)
	for range matrix {
		dp = append(dp, make([]int, len(matrix[0])))
	}
	n1 := len(dp)
	n2 := len(dp[0])

	dp[0][0] = matrix[0][0]
	for i := 1; i < n1; i++ {
		dp[i][0] += matrix[i][0] + dp[i-1][0]
	}
	for j := 1; j < n2; j++ {
		dp[0][j] += matrix[0][j] + dp[0][j-1]
	}

	for i := 1; i < n1; i++ {
		for j := 1; j < n2; j++ {
			dp[i][j] = getMin68(dp[i-1][j], dp[i][j-1]) + matrix[i][j]
		}
	}
	return dp[n1-1][n2-1]
}

func getMin68(a, b int) int {
	if a < b {
		return a
	}
	return b
}
