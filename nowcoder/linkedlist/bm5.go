package linkedlist

import (
	"container/heap"
)

/**
BM5 合并k个已排序的链表
合并 k 个升序的链表并将结果作为一个升序的链表返回其头节点。

数据范围：节点总数 0 \le n \le 50000≤n≤5000，每个节点的val满足 |val| <= 1000∣val∣<=1000
要求：时间复杂度 O(nlogn)O(nlogn)

示例1
输入：
[{1,2,3},{4,5,6,7}]
复制
返回值：
{1,2,3,4,5,6,7}
复制
示例2
输入：
[{1,2},{1,4,5},{6}]
复制
返回值：
{1,1,2,4,5,6}

*/

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int { return len(h) } // 绑定len方法,返回长度
func (h ListNodeHeap) Less(i, j int) bool { // 绑定less方法
	return h[i].Val < h[j].Val // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h ListNodeHeap) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *ListNodeHeap) Pop() any { // 绑定pop方法，从最后拿出一个元素并返回
	if len(*h) <= 0 {
		return nil
	}
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func (h *ListNodeHeap) Push(x any) { // 绑定push方法，插入新元素
	*h = append(*h, x.(*ListNode))
}

func mergeKLists(lists []*ListNode) *ListNode {
	nodeHeap := make(ListNodeHeap, 0, 2*len(lists))
	heap.Init(&nodeHeap)
	newHead := &ListNode{}
	cur := newHead

	for k, _ := range lists {
		if lists[k] != nil {
			heap.Push(&nodeHeap, lists[k])
		}
	}

	for len(nodeHeap) > 0 {
		pop := heap.Pop(&nodeHeap)
		cur.Next = pop.(*ListNode)
		cur = cur.Next
		if cur.Next != nil {
			heap.Push(&nodeHeap, cur.Next)
		}
	}

	return newHead.Next
}


