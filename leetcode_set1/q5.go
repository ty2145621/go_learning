package leetcode_set1

// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
// 输入: "babad"
// 输出: "bab"
// 注意: "aba" 也是一个有效答案。
func longestPalindrome(s string) string {
	var (
		left  = 0
		right = 0
		index = 0
	)
	if len(s) == 0 {
		return ""
	}

	for index < len(s) {
		l := index
		r := index
		for l >= 0 && s[l] == s[index] {
			l--
		}
		for r < len(s) && s[r] == s[index] {
			r++
		}

		index = r

		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if right-left+1 < r-l-1 {
			right = r - 1
			left = l + 1
		}
	}
	return s[left : right+1]
}

/*
func longestPalindrome(s string) string {
	var (
		left    = 0
		right   = 0
		length  = 0
	)
	if len(s) == 0 {
		return ""
	}

	for k, _ := range []byte(s) {
		length = 0
		for k-length >= 0 && k+length < len(s) && s[k+length] == s[k-length] {
			length++
		}
		length--
		if right-left+1 < length*2+1 {
			right = k + length
			left = k - length
		}
		length = 1
		for k-length >= 0 && k-1+length < len(s) && s[k-1+length] == s[k-length] {
			length++
		}
		length--
		if right-left+1 < length*2 {
			right = k + length - 1
			left = k - length
		}
	}
	return s[left : right+1]
}
*/
