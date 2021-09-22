package leetcode_set1

/**
罗马数字转int
 */
func romanToInt(s string) int {
	var m = make(map[uint8]int, 16)
	m = map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return m[s[0]]
	}

	var ret = m[s[len(s)-1]]
	for k := len(s) - 2; k >= 0; k-- {
		if m[s[k]] >= m[s[k+1]] {
			ret += m[s[k]]
		} else {
			ret -= m[s[k]]
		}
	}
	return ret
}
