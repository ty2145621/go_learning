package hash

import (
	"sort"
)

/*
*
BM54 三数之和
中等  通过率：26.81%  时间限制：1秒  空间限制：256M
知识点
数组
双指针
排序
描述
给出一个有n个元素的数组S，S中是否有元素a,b,c满足a+b+c=0？找出数组S中所有满足条件的三元组。

数据范围：0 \le n \le 10000≤n≤1000，数组中各个元素值满足 |val | \le 100∣val∣≤100
空间复杂度：O(n^2)O(n
2

	)，时间复杂度 O(n^2)O(n

2

	)

注意：
三元组（a、b、c）中的元素必须按非降序排列。（即a≤b≤c）
解集中不能包含重复的三元组。
*/
func threeSum(num []int) [][]int {
	// 三元组>>子序列
	// 1.三元组 非降序排列>>将num升序排列
	// 2.数求和问题>>双指针
	res := make([][]int, 0)
	// 排序
	sort.Ints(num)
	n := len(num)
	for i := 0; i < n-2; i++ {
		if num[i] > 0 {
			break // 如果当前数字大于0,则三数之和一定大于0
		}
		if i > 0 && num[i] == num[i-1] {
			continue // 去重
		}
		// 双指针求和
		low, high := i+1, n-1 // 一开始就是最大值和最小值
		for low < high {
			sum := num[i] + num[low] + num[high]
			if sum == 0 {
				// 写入结果
				res = append(res, []int{num[i], num[low], num[high]})
				// 头部去重（如果后面一个数跟当前的数字相等，则代表有重复的结果生成，跳过
				for low < high && num[low] == num[low+1] {
					low++
				}
				low++
				high--
			} else if sum > 0 {
				// 太大前移
				high--
			} else {
				// 太小后移
				low++
			}
		}
	}
	return res
}
