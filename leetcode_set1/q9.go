package leetcode_set1

/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	var (
		y = 0
	)
	for x > y {
		y = y*10 + x%10
		x /= 10
		if x == y || x/10 == y {
			return true
		}
	}

	return false
}
