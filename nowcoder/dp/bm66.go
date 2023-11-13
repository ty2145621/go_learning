package dp

/**
BM66 最长公共子串
动态规划
描述
给定两个字符串str1和str2,输出两个字符串的最长公共子串
题目保证str1和str2的最长公共子串存在且唯一。

数据范围： 1 \le |str1|,|str2| \le 50001≤∣str1∣,∣str2∣≤5000
要求： 空间复杂度 O(n^2)O(n
2
 )，时间复杂度 O(n^2)O(n
2
 )
示例1
输入：
"1AB2345CD","12345EF"
复制
返回值：
"2345"
*/

// LCS66 跟65比起来，要求连续
func LCS66(str1 string, str2 string) string {
	// return LCS2(str1, str2)
	return LCS3(str1, str2)
}

func LCS3(str1 string, str2 string) string {
	// write code here
	var maxLastIndex, maxLength int
	dp := make([]int, len(str2)+1)
	for i := 0; i < len(str1); i++ {
		// 倒序
		for j := len(str2) - 1; j >= 0; j-- {
			if str1[i] == str2[j] {
				dp[j+1] = dp[j] + 1
				if dp[j+1] > maxLength {
					maxLength = dp[j+1]
					maxLastIndex = i
				}
			} else {
				dp[j+1] = 0
			}
		}
	}
	return str1[maxLastIndex-maxLength+1 : maxLastIndex+1]
}

func LCS2(s1 string, s2 string) string {
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

	m := 0
	ri := 0
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if s2[j-1] == s1[i-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
				if m < dp[i][j] {
					m = dp[i][j]
					ri = i
				}
			}
		}
	}
	return s1[ri-m : ri]
}
