package linkedlist

/*
*
BM9 删除链表的倒数第n个节点
题目
题解(224)
讨论(307)
排行
面经new
中等  通过率：40.45%  时间限制：1秒  空间限制：256M
知识点
链表
双指针
描述
给定一个链表，删除链表的倒数第 n 个节点并返回链表的头指针
例如，
给出的链表为: 1\to 2\to 3\to 4\to 51→2→3→4→5, n= 2n=2.
删除了链表的倒数第 nn 个节点之后,链表变为1\to 2\to 3\to 51→2→3→5.

数据范围： 链表长度 0\le n \le 10000≤n≤1000，链表中任意节点的值满足 0 \le val \le 1000≤val≤100
要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
备注：
题目保证 nn 一定是有效的
示例1
输入：
{1,2},2
复制
返回值：
{2}
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{Next: head}
	cur := newHead
	fast := newHead
	n += 1
	for n > 0 && fast != nil {
		n--
		fast = fast.Next
	}

	if n > 0 {
		return newHead.Next
	}

	for fast != nil {
		cur = cur.Next
		fast = fast.Next
	}

	cur.Next = cur.Next.Next
	return newHead.Next
}
