package stack

import (
	"container/list"
)

/**
BM45 滑动窗口的最大值
描述
给定一个长度为 n 的数组 num 和滑动窗口的大小 size ，找出所有滑动窗口里数值的最大值。

例如，如果输入数组{2,3,4,2,6,2,5,1}及滑动窗口的大小3，那么一共存在6个滑动窗口，他们的最大值分别为{4,4,6,6,6,5}； 针对数组{2,3,4,2,6,2,5,1}的滑动窗口有以下6个： {[2,3,4],2,6,2,5,1}， {2,[3,4,2],6,2,5,1}， {2,3,[4,2,6],2,5,1}， {2,3,4,[2,6,2],5,1}， {2,3,4,2,[6,2,5],1}， {2,3,4,2,6,[2,5,1]}。

窗口大于数组长度或窗口长度为0的时候，返回空。

数据范围： 1 \le n \le 100001≤n≤10000，0 \le size \le 100000≤size≤10000，数组中每个元素的值满足 |val| \le 10000∣val∣≤10000
要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
*/

func maxInWindows(num []int, size int) []int {
	if size == 0 || size > len(num) {
		return nil
	}
	s := list.New()
	res := make([]int, 0)
	for k := 0; k < size; k++ {
		for s.Len() != 0 && num[s.Back().Value.(int)] < num[k] {
			s.Remove(s.Back())
		}
		s.PushBack(k)
	}
	for k := size; k < len(num); k++ {
		res = append(res, num[s.Front().Value.(int)])
		for s.Len() != 0 && s.Front().Value.(int) < (k-size+1) {
			s.Remove(s.Front())
		}
		for s.Len() != 0 && num[s.Back().Value.(int)] < num[k] {
			s.Remove(s.Back())
		}
		s.PushBack(k)
	}
	res = append(res, num[s.Front().Value.(int)])
	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
