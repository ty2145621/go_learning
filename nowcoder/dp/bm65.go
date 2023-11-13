package dp

/**
BM65 最长公共子序列(二)
题目
题解(105)
讨论(131)
排行
面经new
中等  通过率：38.12%  时间限制：3秒  空间限制：256M
知识点
动态规划
描述
给定两个字符串str1和str2，输出两个字符串的最长公共子序列。如果最长公共子序列为空，则返回"-1"。目前给出的数据，仅仅会存在一个最长的公共子序列

数据范围：0 \le |str1|,|str2| \le 20000≤∣str1∣,∣str2∣≤2000
要求：空间复杂度 O(n^2)O(n
2
 ) ，时间复杂度 O(n^2)O(n
2
 )
*/

func LCS(s1 string, s2 string) string {
	return LCS1(s1, s2)
}

func LCS1(s1 string, s2 string) string {
	n1 := len(s1)
	n2 := len(s2)
	if n1 == 0 || n2 == 0 {
		return ""
	}

	dp := make([][]int, 0)
	for range []byte(s1) {
		dp = append(dp, make([]int, n2+1))
	}
	dp = append(dp, make([]int, n2+1))

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if s2[j-1] == s1[i-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = getMax65(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	res := make([]byte, 0)
	for i, j := n1, n2; dp[i][j] > 0; {
		if s1[i-1] == s2[j-1] {
			res = append(res, s1[i-1])
			j--
			i--
		} else if dp[i-1][j] == dp[i][j] {
			i--
		} else {
			j--
		}
	}

	if len(res) == 0 {
		return "-1"
	}
	res2 := make([]byte, 0)
	for i := len(res) - 1; i >= 0; i-- {
		res2 = append(res2, res[i])
	}
	return string(res2)
}

func getMax65(a ...int) int {
	m := a[0]
	for _, v := range a {
		if m < v {
			m = v
		}
	}
	return m
}
