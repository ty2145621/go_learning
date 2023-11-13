package dp

/**
BM69 把数字翻译成字符串
题目
题解(95)
讨论(97)
排行
面经new
中等  通过率：24.21%  时间限制：1秒  空间限制：256M
知识点
动态规划
描述
有一种将字母编码成数字的方式：'a'->1, 'b->2', ... , 'z->26'。

现在给一串数字，返回有多少种可能的译码结果

数据范围：字符串长度满足 0 < n \le 900<n≤90
进阶：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
*/

func solve(nums string) int {
	return solve1(nums)
}
func solve1(nums string) int {
	if len(nums) == 0 {
		return 0
	}
	if nums[len(nums)-1] == '0' && (len(nums) == 1 || nums[len(nums)-2] > '2') {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	// dp i+1 表示 nums i 位置
	dp := make([]int, len(nums)+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < len(nums); i++ {
		if nums[i-1] == '0' {
			if nums[i] != '0' {
				dp[i+1] = dp[i]
			} else {
				return 0
			}
		} else if nums[i-1] == '1' || nums[i-1] == '2' {
			if nums[i] != '0' {
				if nums[i-1] == '2' && nums[i] >= '7' {
					dp[i+1] = dp[i]
				} else {
					dp[i+1] = dp[i] + dp[i-1]
				}
			} else {
				dp[i+1] = dp[i-1]
			}
		} else {
			dp[i+1] = dp[i]
		}
	}

	return dp[len(nums)]
}
