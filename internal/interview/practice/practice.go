package practice

import (
	"container/heap"
	"strconv"
)

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	l, r := 0, len(arr)-1
	p := arr[0]
	for l < r {
		for l < r && arr[r] >= p {
			r--
		}
		arr[l] = arr[r]
		for l < r && arr[l] <= p {
			l++
		}
		arr[r] = arr[l]
	}
	arr[l] = p
	quickSort(arr[:l])
	quickSort(arr[l+1:])
}

func quickSort2(arr []int) {
	if len(arr) <= 1 {
		return
	}
	l, r := 0, len(arr)-1
	p := arr[0]
	for l < r {
		for l < r && arr[r] >= p {
			r--
		}
		arr[l] = arr[r]
		for l < r && arr[l] <= p {
			l++
		}
		arr[r] = arr[l]
	}
	arr[l] = p
	quickSort2(arr[:l])
	quickSort2(arr[l+1:])
}

type listNode struct {
	n    int
	next *listNode
}

func (l *listNode) String() string {
	t := l
	s := ""
	for t != nil {
		s += strconv.Itoa(t.n) + ","
		t = t.next
	}
	return s
}

func (l *listNode) build(iArr []int) {
	t := l
	for k, i := range iArr {
		t.n = i
		if k != len(iArr)-1 {
			t.next = &listNode{}
			t = t.next
		}
	}
}

func reverseList(head *listNode) *listNode {
	var pre *listNode
	cur := head

	for cur != nil {
		cur, pre, cur.next = cur.next, cur, pre
	}
	return pre
}

func reverseKList(head *listNode, k int) *listNode {
	if k <= 1 {
		return head
	}
	newHead := &listNode{next: head}
	pre := newHead
	cur := head
	for cur != nil {
		if !checkCanReverse(cur, k) {
			break
		}
		pre.next = reverseKListPartition(cur, k)
		pre, cur = cur, cur.next
	}

	return newHead.next
}

func reverseKList2(head *listNode, k int) *listNode {
	if k <= 1 {
		return head
	}
	newHead := &listNode{next: head}
	pre := newHead
	cur := head

	for cur != nil {
		pre.next = reverseKListPartition(cur, k)
		if cur == nil {
			break
		}
		pre, cur = cur, cur.next
	}

	return newHead.next
}

func checkCanReverse(head *listNode, k int) bool {
	cur := head
	for i := 0; i < k; i++ {
		if cur == nil {
			return false
		}
		cur = cur.next
	}
	return true
}

func reverseKListPartition(head *listNode, k int) *listNode {
	var pre *listNode
	cur := head
	for cur != nil && k > 0 {
		pre, cur, cur.next = cur, cur.next, pre
		k--
	}

	head.next = cur

	return pre
}

func search[T ~int](arr []T, target T) int {
	l, r := 0, len(arr)-1
	mid := (l + r) / 2

	for l <= r {
		if target == arr[mid] {
			return mid
		}
		if target > arr[mid] {
			l = mid + 1
		} else if target < arr[mid] {
			r = mid - 1
		}
		mid = (l + r) / 2
	}
	return -1
}

type heapInt []*listNode

// Len is the number of elements in the collection.
func (h heapInt) Len() int {
	return len(h)
}

func (h heapInt) Less(i int, j int) bool {
	return h[i].n < h[j].n
}

// Swap swaps the elements with indexes i and j.
func (h heapInt) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heapInt) Push(x any) {
	*h = append(*h, x.(*listNode))
}

func (h *heapInt) Pop() any {
	len := h.Len()
	t := (*h)[len-1]
	*h = (*h)[:len-1]
	return any(t)
}

func mergeKLists(l []*listNode) *listNode {
	minHeap := make(heapInt, 0, len(l))
	heap.Init(&minHeap)

	for _, list := range l {
		heap.Push(&minHeap, list)
	}

	newHead := &listNode{}
	cur := newHead
	for minHeap.Len() > 0 {
		tmp := heap.Pop(&minHeap).(*listNode)
		cur.next = tmp
		if tmp.next != nil {
			heap.Push(&minHeap, tmp.next)
		}
		cur = cur.next
	}

	return newHead.next
}


