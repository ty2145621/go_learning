package twopointer

/**
BM91 反转字符串
题目
题解(175)
讨论(299)
排行
面经new
入门  通过率：67.00%  时间限制：1秒  空间限制：256M
知识点
双指针
字符串
描述
写出一个程序，接受一个字符串，然后输出该字符串反转后的字符串。（字符串长度不超过1000）

数据范围： 0 \le n \le 10000≤n≤1000
要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
*/

func solve91(str string) string {
	bArr := make([]byte, len(str))
	for k, v := range []byte(str) {
		bArr[len(bArr)-k-1] = v
	}
	return string(bArr)
}