package stack

/**
BM47 寻找第K大
中等  通过率：29.45%  时间限制：1秒  空间限制：256M
知识点
堆
分治
描述
有一个整数数组，请你根据快速排序的思路，找出数组中第 k 大的数。

给定一个整数数组 a ,同时给定它的大小n和要找的 k ，请返回第 k 大的数(包括重复的元素，不用去重)，保证答案存在。
要求：时间复杂度 O(nlogn)O(nlogn)，空间复杂度 O(1)O(1)
数据范围：0\le n \le 10000≤n≤1000， 1 \le K \le n1≤K≤n，数组中每个元素满足 0 \le val \le 100000000≤val≤10000000
示例1
输入：
[1,3,5,2,2],5,3
复制
返回值：
2
复制
示例2
输入：
[10,10,9,9,8,7,5,6,4,3,4,2],12,3
复制
返回值：
9
复制
说明：
去重后的第3大是8，但本题要求包含重复的元素，不用去重，所以输出9
*/

func findKth(a []int, n int, K int) int {
	a = a[0:n]
	p := a[0]
	left := 1
	for j := left; j < n; j++ {
		if a[j] <= p {
			continue
		}
		a[left], a[j] = a[j], a[left]
		left++
	}
	left--
	a[left], a[0] = a[0], a[left]
	// j 的位置
	if left+1 == K {
		return a[left]
	}
	if left+1 > K {
		return findKth(a[0:left], left, K)
	} else {
		return findKth(a[left+1:], n-left-1, K-left-1)
	}
}
