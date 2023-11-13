package real

/*
*
BM97 旋转数组
面经new
中等  通过率：49.46%  时间限制：1秒  空间限制：256M
知识点
数组
描述
一个数组A中存有 n 个整数，在不允许使用另外数组的前提下，将每个整数循环向右移 M（ M >=0）个位置，
即将A中的数据由（A0 A1 ……AN-1 ）变换为（AN-M …… AN-1 A0 A1 ……AN-M-1 ）（最后 M 个数循环移至最前面的 M 个位置）。
如果需要考虑程序移动数据的次数尽量少，要如何设计移动的方法？

输入：
6,2,[1,2,3,4,5,6]
复制
返回值：
[5,6,1,2,3,4]
*/
func solve(n int, m int, a []int) []int {
	if m > n {
		m %= n
	}
	if len(a) <= 1 {
		return a
	}
	reverse(a, 0, n-m-1)
	reverse(a, n-m, n-1)
	reverse(a, 0, n-1)
	return a
}

func reverse(a []int, l, r int) {
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}
