package leetcode_set1

// 3. 无重复字符的最长子串

//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
func lengthOfLongestSubstring(s string) int {
	var m = make(map[byte]int)
	var longest = 0
	var begin = 0 // 字符串开始的位置
	for k, v := range []byte(s) {
		if l, ok := m[v]; ok {
			if l >= begin {
				if k-begin > longest {
					longest = k - begin
				}
				begin = l + 1
			}
		}
		m[v] = k
	}
	if len(s)-begin > longest {
		longest = len(s) - begin
	}
	return longest
}

/*func lengthOfLongestSubstring(s string) int {
	var freq [256]int
	result, left, right := 0, 0, 0
	for left < len(s) {
		if right < len(s) && freq[s[right]] == 0 {
			freq[s[right]] ++
			right++
		} else {
			freq[s[left]]--
			left++
		}
		result = max(result, right-left)
	}
	return result
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}*/
