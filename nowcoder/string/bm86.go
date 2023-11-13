package string

/**
BM86 大数加法
题目
题解(255)
讨论(333)
排行
面经new
中等  通过率：43.61%  时间限制：1秒  空间限制：256M
知识点
字符串
模拟
描述
以字符串的形式读入两个数字，编写一个函数计算它们的和，以字符串形式返回。

数据范围：s.length,t.length \le 100000s.length,t.length≤100000，字符串仅由'0'~‘9’构成
要求：时间复杂度 O(n)O(n)
*/

func solve86(s string, t string) string {
	p := 0
	jinwei := byte(0)
	res := make([]byte, 0)
	for p < len(s) && p < len(t) {
		c := byte(s[len(s)-p-1]) + byte(t[len(t)-p-1]) - '0' - '0' + jinwei
		num := c % 10
		jinwei = c / 10
		res = append(res, num+'0')
		p++
	}
	for p < len(s) {
		c := byte(s[len(s)-p-1]) - '0' + jinwei
		num := c % 10
		jinwei = c / 10
		res = append(res, num+'0')
		p++
	}
	for p < len(t) {
		c := byte(t[len(t)-p-1]) - '0' + jinwei
		num := c % 10
		jinwei = c / 10
		res = append(res, num+'0')
		p++
	}
	for jinwei >= 1 {
		c := jinwei
		num := c % 10
		jinwei = c / 10
		res = append(res, num+'0')
	}

	for k := 0; k < len(res)/2; k++ {
		res[k], res[len(res)-k-1] = res[len(res)-k-1], res[k]
	}
	return string(res)
}
