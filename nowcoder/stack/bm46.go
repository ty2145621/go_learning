package stack

import (
	"container/heap"
)

type IntMaxHeap []int // 定义一个类型

func (h IntMaxHeap) Len() int { return len(h) } // 绑定len方法,返回长度
func (h IntMaxHeap) Less(i, j int) bool { // 绑定less方法
	return h[i] > h[j] // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h IntMaxHeap) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *IntMaxHeap) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntMaxHeap) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Top() int {
	// return (*h)[len(*h)/2-1]
	return (*h)[0]
}

func GetLeastNumbers_Solution(input []int, k int) []int {
	return GetLeastNumbers_Solution1(input, k)
}

func GetLeastNumbers_Solution1(input []int, k int) []int {
	if k == 0 {
		return nil
	}
	if k >= len(input) {
		return input
	}

	pq := IntMaxHeap{}
	heap.Init(&pq)

	for _, v := range input {
		if pq.Len() < k {
			heap.Push(&pq, v)
		} else {
			top := pq.Top()
			if top > v {
				pq[0] = v
				heap.Fix(&pq, 0)
			}
		}
	}
	res := make([]int, 0)
	for pq.Len() > 0 {
		res = append(res, heap.Pop(&pq).(int))
	}
	return res
}
