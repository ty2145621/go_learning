package linkedlist

/**
BM15 删除有序链表中重复的元素-I

删除给出链表中的重复元素（链表中元素从小到大有序），使链表中的所有元素都只出现一次
例如：
给出的链表为1\to1\to21→1→2,返回1 \to 21→2.
给出的链表为1\to1\to 2 \to 3 \to 31→1→2→3→3,返回1\to 2 \to 31→2→3.

数据范围：链表长度满足 0 \le n \le 1000≤n≤100，链表中任意节点的值满足 |val| \le 100∣val∣≤100
进阶：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
示例1
输入：
{1,1,2}
复制
返回值：
{1,2}
复制
示例2
输入：
{}
复制
返回值：
{}
*/

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head

	var next *ListNode
	for cur != nil && cur.Next != nil {
		next = cur.Next
		for next != nil && next.Val == cur.Val {
			next = next.Next
		}
		cur, cur.Next = next, next
	}

	return head
}
