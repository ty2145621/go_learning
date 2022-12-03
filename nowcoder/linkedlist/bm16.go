package linkedlist

/*
*
BM16 删除有序链表中重复的元素-II
给出一个升序排序的链表，删除链表中的所有重复出现的元素，只保留原链表中只出现一次的元素。
例如：
给出的链表为1 \to 2\to 3\to 3\to 4\to 4\to51→2→3→3→4→4→5, 返回1\to 2\to51→2→5.
给出的链表为1\to1 \to 1\to 2 \to 31→1→1→2→3, 返回2\to 32→3.

数据范围：链表长度 0 \le n \le 100000≤n≤10000，链表中的值满足 |val| \le 1000∣val∣≤1000
要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
进阶：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
示例1
输入：
{1,2,2}
复制
返回值：
{1}
复制
示例2
输入：
{}
复制
返回值：
{}
*/
func deleteDuplicates2(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}

	prev := newHead
	cur := newHead.Next

	var next *ListNode
	for cur != nil && cur.Next != nil {
		next = cur.Next
		for next != nil && next.Val == cur.Val {
			next = next.Next
		}

		if cur.Next != next {
			prev.Next, cur = next, next
		} else {
			prev, cur = cur, next
		}
	}
	return newHead.Next
}
