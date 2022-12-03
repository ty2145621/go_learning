package linkedlist

/**
BM2 链表内指定区间反转
描述
将一个节点数为 size 链表 m 位置到 n 位置之间的区间反转，要求时间复杂度 O(n)O(n)，空间复杂度 O(1)O(1)。
例如：
给出的链表为 1\to 2 \to 3 \to 4 \to 5 \to NULL1→2→3→4→5→NULL, m=2,n=4m=2,n=4,
返回 1\to 4\to 3\to 2\to 5\to NULL1→4→3→2→5→NULL.

数据范围： 链表长度 0 < size \le 10000<size≤1000，0 < m \le n \le size0<m≤n≤size，链表中每个节点的值满足 |val| \le 1000∣val∣≤1000
要求：时间复杂度 O(n)O(n) ，空间复杂度 O(n)O(n)
进阶：时间复杂度 O(n)O(n)，空间复杂度 O(1)O(1)
示例1
输入：
{1,2,3,4,5},2,4
复制
返回值：
{1,4,3,2,5}
复制
示例2
输入：
{5},1,1
复制
返回值：
{5}

*/

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	newHead := &ListNode{Next: head}
	prev := newHead
	for i := 1; i < m; i++ {
		prev = prev.Next
	}

	prev.Next = reverse1(prev.Next, m, n)

	return newHead.Next
}

func reverse1(head *ListNode, m int, n int) *ListNode {
	cur := head
	var prev *ListNode

	for i := m; i <= n; i++ {
		prev, cur, cur.Next = cur, cur.Next, prev
	}

	head.Next = cur
	return prev
}
