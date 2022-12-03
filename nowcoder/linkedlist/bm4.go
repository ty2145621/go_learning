package linkedlist

/**
输入两个递增的链表，单个链表的长度为n，合并这两个链表并使新链表中的节点仍然是递增排序的。
数据范围： 0 \le n \le 10000≤n≤1000，-1000 \le 节点值 \le 1000−1000≤节点值≤1000
要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)

如输入{1,3,5},{2,4,6}时，合并后的链表为{1,2,3,4,5,6}，所以对应的输出为{1,2,3,4,5,6}，转换过程如下图所示：

或输入{-1,2,4},{1,3,4}时，合并后的链表为{-1,1,2,3,4,4}，所以对应的输出为{-1,1,2,3,4,4}，转换过程如下图所示：

示例1
输入：
{1,3,5},{2,4,6}
复制
返回值：
{1,2,3,4,5,6}
复制
示例2
输入：
{},{}
复制
返回值：
{}
复制
示例3
输入：
{-1,2,4},{1,3,4}
复制
返回值：
{-1,1,2,3,4,4}
复制
相似企业真题

*/

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
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
