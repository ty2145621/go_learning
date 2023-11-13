package greedy

import (
	"container/heap"
	"sort"
)

/**
BM96 主持人调度（二）
描述
有 n 个活动即将举办，每个活动都有开始时间与活动的结束时间，第 i 个活动的开始时间是 starti ,第 i 个活动的结束时间是 endi ,举办某个活动就需要为该活动准备一个活动主持人。

一位活动主持人在同一时间只能参与一个活动。并且活动主持人需要全程参与活动，换句话说，一个主持人参与了第 i 个活动，那么该主持人在 (starti,endi) 这个时间段不能参与其他任何活动。求为了成功举办这 n 个活动，最少需要多少名主持人。

数据范围: 1 \le n \le 10^51≤n≤10
5
  ， -2^{32} \le start_i\le end_i \le 2^{31}-1−2
32
 ≤start
i
​
 ≤end
i
​
 ≤2
31
 −1

复杂度要求：时间复杂度 O(n \log n)O(nlogn) ，空间复杂度 O(n)O(n)
*/

type IntMinHeap []int // 定义一个类型

func (h IntMinHeap) Len() int { return len(h) } // 绑定len方法,返回长度
func (h IntMinHeap) Less(i, j int) bool { // 绑定less方法
	return h[i] < h[j] // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h IntMinHeap) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *IntMinHeap) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntMinHeap) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Top() int { // 绑定push方法，插入新元素
	return (*h)[0]
}

func minmumNumberOfHost(n int, startEnd [][]int) int {
	sort.Slice(startEnd, func(i, j int) bool {
		return startEnd[i][0] < startEnd[j][0]
	})

	arr := make(IntMinHeap, 0)
	heap.Init(&arr)

	max := 0
	for _, v := range startEnd {
		for arr.Len() > 0 && arr.Top() <= v[0] {
			heap.Pop(&arr)
		}
		heap.Push(&arr, v[1])
		if arr.Len() > max {
			max = arr.Len()
		}
	}
	return max
}

func minmumNumberOfHost2(n int, startEnd [][]int) int {
	starts := make([]int, len(startEnd))
	ends := make([]int, len(startEnd))
	for i := 0; i < len(startEnd); i++ {
		starts[i] = startEnd[i][0]
		ends[i] = startEnd[i][1]
	}
	sort.Ints(starts)
	sort.Ints(ends)

	e := 0
	count := 0
	for i := range startEnd {
		if starts[i] < ends[e] {
			count++
		} else {
			e++
		}
	}

	return count
}

func minmumNumberOfHost3(n int, startEnd [][]int) int {
	points := make([]int, 0, len(startEnd)*2)
	m := make(map[int]int, 0)
	for i := 0; i < len(startEnd); i++ {
		if _, ok := m[startEnd[i][0]]; !ok {
			points = append(points, startEnd[i][0])
		}
		m[startEnd[i][0]]++
		if _, ok := m[startEnd[i][1]]; !ok {
			points = append(points, startEnd[i][1])
		}
		m[startEnd[i][1]]--
	}
	sort.Ints(points)

	sum, max := 0, 0
	for _, v := range points {
		sum += m[v]
		if max < sum {
			max = sum
		}
	}

	return max
}
