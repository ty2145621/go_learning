package dp

/**
BM67 不同路径的数目(一)
题目
题解(113)
讨论(224)
排行
面经new
简单  通过率：49.52%  时间限制：1秒  空间限制：256M
知识点
动态规划
描述
一个机器人在m×n大小的地图的左上角（起点）。
机器人每次可以向下或向右移动。机器人要到达地图的右下角（终点）。
可以有多少种不同的路径从起点走到终点？
*/

func uniquePaths(m int, n int) int {
	// write code here
	return uniquePaths1(m, n)
}

func uniquePaths1(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	dp := make([][]int, 0)
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]int, n+1))
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
