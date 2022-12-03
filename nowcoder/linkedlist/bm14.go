package linkedlist

/**
BM14 链表的奇偶重排
给定一个单链表，请设定一个函数，将链表的奇数位节点和偶数位节点分别放在一起，重排后输出。
注意是节点的编号而非节点的数值。

数据范围：节点数量满足 0 \le n \le 10^50≤n≤10
5
 ，节点中的值都满足 0 \le val \le 10000≤val≤1000
要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
示例1
输入：
{1,2,3,4,5,6}
复制
返回值：
{1,3,5,2,4,6}
复制
说明：
1->2->3->4->5->6->NULL
重排后为
1->3->5->2->4->6->NULL

*/

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	odd := head
	even := head.Next
	evenFirst := even

	for even != nil && even.Next != nil {
		odd, even, odd.Next, even.Next = even.Next, even.Next.Next, even.Next, even.Next.Next
	}
	odd.Next = evenFirst

	return head
}
