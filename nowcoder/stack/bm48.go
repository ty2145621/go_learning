package stack

import (
	"container/heap"
)

/**
BM48 数据流中的中位数
中等  通过率：29.80%  时间限制：1秒  空间限制：256M
知识点
排序
堆
描述
如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。我们使用Insert()方法读取数据流，使用GetMedian()方法获取当前读取数据的中位数。

数据范围：数据流中数个数满足 1 \le n \le 1000 \1≤n≤1000  ，大小满足 1 \le val \le 1000 \1≤val≤1000

进阶： 空间复杂度 O(n) \O(n)  ， 时间复杂度 O(nlogn) \O(nlogn)
示例1
输入：
[5,2,3,4,1,6,7,0,8]
复制
返回值：
"5.00 3.50 3.00 3.50 3.00 3.50 4.00 3.50 4.00 "
复制
说明：
数据流里面不断吐出的是5,2,3...,则得到的平均数分别为5,(5+2)/2,3...
示例2
输入：
[1,1,1]
复制
返回值：
"1.00 1.00 1.00 "
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

func (h *IntMinHeap) Top() int {
	// return (*h)[len(*h)/2-1]
	return (*h)[0]
}

var pq481 = IntMaxHeap{}
var pq482 = IntMinHeap{}

func init() {
	heap.Init(&pq481)
	heap.Init(&pq482)
}

func Insert(num int) {
	if float64(num) < GetMedian() {
		heap.Push(&pq481, num)
	} else {
		heap.Push(&pq482, num)
	}
	for pq481.Len() > pq482.Len() {
		heap.Push(&pq482, heap.Pop(&pq481))
	}
	for pq481.Len() < pq482.Len() {
		heap.Push(&pq481, heap.Pop(&pq482))
	}
}

func GetMedian() float64 {
	if pq481.Len() == 0 {
		return 0
	}
	if pq481.Len() == pq482.Len() {
		return float64(pq481.Top()+pq482.Top()) / 2
	}
	if pq481.Len() > pq482.Len() {
		return float64(pq481.Top())
	}
	if pq481.Len() < pq482.Len() {
		return float64(pq482.Top())
	}
	return 0
}
