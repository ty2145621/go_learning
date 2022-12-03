package linkedlist

/**
BM12 单链表的排序
给定一个节点数为n的无序单链表，对其按升序排序。

数据范围：0 < n \le 1000000<n≤100000
要求：空间复杂度 O(n)O(n)，时间复杂度 O(nlogn)O(nlogn)

示例1
输入：
[1,3,2,4,5]
复制
返回值：
{1,2,3,4,5}
复制
示例2
输入：
[-1,0,-2]
复制
返回值：
{-2,-1,0}
复制
*/

func sortInList(head *ListNode) *ListNode {
	return sort1(head)
}

func sort1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Next.Next == nil {
		if head.Next.Val < head.Val {
			head.Val, head.Next.Val = head.Next.Val, head.Val
		}
		return head
	}

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 切断slow节点的next
	head2 := slow.Next
	slow.Next = nil

	h1 := sort1(head)
	h2 := sort1(head2)
	return merge2List(h1, h2)
}

func merge2List(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	newHead := &ListNode{}
	h1 := pHead1
	h2 := pHead2
	cur := newHead

	for h1 != nil && h2 != nil {
		if h1.Val <= h2.Val {
			cur.Next = h1
			h1 = h1.Next
			cur = cur.Next
		} else {
			cur.Next = h2
			h2 = h2.Next
			cur = cur.Next
		}
	}

	for h1 != nil {
		cur.Next = h1
		h1 = h1.Next
		cur = cur.Next
	}
	for h2 != nil {
		cur.Next = h2
		h2 = h2.Next
		cur = cur.Next
	}

	return newHead.Next
}
