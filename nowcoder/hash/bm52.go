package hash

/**

BM52 数组中只出现一次的两个数字
题目
题解(164)
讨论(186)
排行
面经new
中等  通过率：58.20%  时间限制：1秒  空间限制：256M
知识点
位运算
哈希
描述
一个整型数组里除了两个数字只出现一次，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字。

数据范围：数组长度 2\le n \le 10002≤n≤1000，数组中每个数的大小 0 < val \le 10000000<val≤1000000
要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)

提示：输出时按非降序排列。
示例1
*/

func FindNumsAppearOnce(array []int) []int {
	a := 0
	for _, v := range array {
		a ^= v
	}

	mask := 1
	for a&mask == 0 {
		mask <<= 1
	}

	p := 0
	q := 0
	for _, v := range array {
		if v&mask == 0 {
			p ^= v
		} else {
			q ^= v
		}
	}

	if p > q {
		return []int{q, p}
	}
	return []int{p, q}
}
