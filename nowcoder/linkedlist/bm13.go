package linkedlist

/**
BM13 判断一个链表是否为回文结构
给定一个链表，请判断该链表是否为回文结构。
回文是指该字符串正序逆序完全一致。
数据范围： 链表节点数 0 \le n \le 10^50≤n≤10
5
 ，链表中每个节点的值满足 |val| \le 10^7∣val∣≤10
7

示例1
输入：
{1}
复制
返回值：
true
复制
示例2
输入：
{2,1}
复制
返回值：
false
复制
说明：
2->1
示例3
输入：
{1,2,2,1}
复制
返回值：
true
复制
说明：
1->2->2->1
*/

func isPail(head *ListNode) bool {
	// 只反转一半
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	newHead := reverse13(slow)

	h1 := head
	h2 := newHead
	for h1 != nil && h2 != nil {
		if h1.Val != h2.Val {
			return false
		}
		h1, h2 = h1.Next, h2.Next
	}

	return true
}

func reverse13(head *ListNode) *ListNode {
	var cur *ListNode
	next := head

	for next != nil {
		cur, next, next.Next = next, next.Next, cur
	}

	return cur
}
