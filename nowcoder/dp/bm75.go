package dp

/**
BM75 编辑距离(一)
题目
题解(38)
讨论(23)
排行
面经new
较难  通过率：56.59%  时间限制：1秒  空间限制：256M
知识点
动态规划
字符串
描述
给定两个字符串 str1 和 str2 ，请你算出将 str1 转为 str2 的最少操作数。
你可以对字符串进行3种操作：
1.插入一个字符
2.删除一个字符
3.修改一个字符。

字符串长度满足 1 \le n \le 1000 \1≤n≤1000  ，保证字符串中只出现小写英文字母。
*/

func editDistance(str1 string, str2 string) int {
	// write code here
	return distance1(str1, str2)
}

func distance1(str1, str2 string) int {
	dp := make([][]int, 0)
	for range []byte(str1) {
		dp = append(dp, make([]int, len(str2)+1))
	}
	dp = append(dp, make([]int, len(str2)+1))

	dp[0][0] = 0
	for i := 1; i <= len(str1); i++ {
		dp[i][0] = dp[i-1][0] + 1
	}
	for j := 1; j <= len(str2); j++ {
		dp[0][j] = dp[0][j-1] + 1
	}

	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = getMin75(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
			}
		}
	}
	return dp[len(str1)][len(str2)]
}

func getMin75(a ...int) int {
	m := a[0]
	for _, v := range a {
		if m > v {
			m = v
		}
	}
	return m
}
