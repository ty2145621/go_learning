package hash

/**
BM53 缺失的第一个正整数
题目
题解(90)
讨论(81)
排行
面经new
中等  通过率：45.20%  时间限制：1秒  空间限制：256M
知识点
数组
哈希
描述
给定一个无重复元素的整数数组nums，请你找出其中没有出现的最小的正整数

进阶： 空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)

数据范围:
-2^{31}\le nums[i] \le 2^{31}-1−2
31
 ≤nums[i]≤2
31
 −1
0\le len(nums)\le5*10^50≤len(nums)≤5∗10
5
*/

func minNumberDisappeared(nums []int) int {
	res := make([]int, len(nums))
	for _, v := range nums {
		if v < len(nums) && v >= 0 {
			res[v] = 1
		}
	}
	for k, v := range res {
		if k != 0 && v != 1 {
			return k
		}
	}
	return len(nums) + 1
}
