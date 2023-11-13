package twopointer

import (
	"sort"
)

/**
BM89 合并区间
给出一组区间，请合并所有重叠的区间。
请保证合并后的区间按区间起点升序排列。

数据范围：区间组数 0 \le n \le 2 \times 10^50≤n≤2×10
5
 ，区间内 的值都满足 0 \le val \le 2 \times 10^50≤val≤2×10
5

要求：空间复杂度 O(n)O(n)，时间复杂度 O(nlogn)O(nlogn)
进阶：空间复杂度 O(val)O(val)，时间复杂度O(val)O(val)
*/

type Interval struct {
	Start int
	End   int
}

/**
 *
 * @param intervals Interval类一维数组
 * @return Interval类一维数组
 */
func merge(intervals []*Interval) []*Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	start, end := intervals[0].Start, intervals[0].End
	res := make([]*Interval, 0)
	for k := 1; k < len(intervals); k++ {
		if intervals[k].Start > end {
			res = append(res, &Interval{Start: start, End: end})
			start, end = intervals[k].Start, intervals[k].End
		} else {
			end = getMax89(end, intervals[k].End)
		}
	}
	res = append(res, &Interval{Start: start, End: end})
	return res
}

func getMax89(a ...int) int {
	m := a[0]
	for _, v := range a {
		if m < v {
			m = v
		}
	}
	return m
}
