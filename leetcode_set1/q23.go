package code_set1

import "container/heap"

/**
23. 合并K个升序链表
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。



示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNodeMinHeap []*ListNode // 定义一个类型

func (h ListNodeMinHeap) Len() int { return len(h) } // 绑定len方法,返回长度
func (h ListNodeMinHeap) Less(i, j int) bool { // 绑定less方法
    return h[i].Val < h[j].Val // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h ListNodeMinHeap) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
    h[i], h[j] = h[j], h[i]
}

func (h *ListNodeMinHeap) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func (h *ListNodeMinHeap) Push(x interface{}) { // 绑定push方法，插入新元素
    *h = append(*h, x.(*ListNode))
}

func mergeKLists(lists []*ListNode) *ListNode {
    var minHeap ListNodeMinHeap = make(ListNodeMinHeap, 0)
    var newHead *ListNode = &ListNode{
        Val:  0,
        Next: nil,
    }
    var cPtr = newHead

    for _, v := range lists {
        if v != nil {
            minHeap.Push(v)
        }
    }
    if len(minHeap) == 0 {
        return newHead.Next
    }
    heap.Init(&minHeap)

    for minHeap.Len() > 0 {
        l := heap.Pop(&minHeap).(*ListNode)
        if l.Next != nil {
            heap.Push(&minHeap, l.Next)
        }
        cPtr.Next = l
        cPtr = cPtr.Next
    }
    cPtr.Next = nil

    return newHead.Next
}
