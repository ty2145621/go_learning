package code_set1

//将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
// 输入: s = "LEETCODEISHIRING", numRows = 4
//输出: "LDREOEIIECIHNTSG"
//解释:
//
//L     D     R
//E   O E   I I
//E C   I H   N
//T     S     G
//
func convert(s string, numRows int) string {
	strRet := []byte(s)
	circle := numRows*2 - 2
	if circle <= 0 {
		return s
	}
	var pos = 0
	for i := 0; i < numRows; i += 1 {
		next := i
		for j := 0; next < len(s); j++ {
			strRet[pos] = s[next]
			pos++
			if i == 0 || i == numRows-1 {
				next += circle
			} else {
				if j%2 == 0 {
					next += circle - 2*i
				} else {
					next += 2 * i
				}
			}
		}
	}
	return string(strRet)
}

/*func convert(s string, numRows int) string {
	strRet := ""
	circle := numRows*2 - 2
    if circle <= 0 {
		return s
	}
	for i := 0; i < numRows; i += 1 {
		next := i
		for j := 0; next < len(s); j++ {
			if i == 0 || i == numRows-1 {
				strRet += fmt.Sprintf("%c%s", s[next], strings.Repeat(" ", numRows-2))
				next += circle
			} else {
				if j%2 == 0 {
					strRet += fmt.Sprintf("%c%s", s[next], strings.Repeat(" ", numRows-2-i))
					next += circle - 2*i
				} else {
					strRet += fmt.Sprintf("%c%s", s[next], strings.Repeat(" ", i-1))
					next += 2 * i
				}
			}
		}
		strRet = strings.TrimRight(strRet, " ")
		strRet += fmt.Sprintf("\n")
	}
	return strRet
}*/
