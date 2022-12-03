package linkedlist

/**

BM11 链表相加(二)
描述
假设链表中每一个节点的值都在 0 - 9 之间，那么链表整体就可以代表一个整数。
给定两个这种链表，请生成代表两个整数相加值的结果链表。
数据范围：0 \le n,m \le 10000000≤n,m≤1000000，链表任意值 0 \le val \le 90≤val≤9
要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)

例如：链表 1 为 9->3->7，链表 2 为 6->3，最后生成新的结果链表为 1->0->0->0。

示例1
输入：
[9,3,7],[6,3]
返回值：
{1,0,0,0}
说明：
如题面解释
示例2
输入：
[0],[6,3]
复制
返回值：
{6,3}
复制
*/

func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
	// write code here
	r1 := reverse11(head1)
	r2 := reverse11(head2)

	c1 := r1
	c2 := r2

	var resHead = &ListNode{Val: 0, Next: nil}
	var resLast = resHead

	var jinwei = 0
	for c1 != nil && c2 != nil {
		newVal := c1.Val + c2.Val + jinwei
		jinwei = newVal / 10
		newNode := &ListNode{Val: newVal % 10, Next: nil}
		resLast.Next = newNode
		resLast = newNode
		c1 = c1.Next
		c2 = c2.Next
	}

	for c1 != nil {
		newVal := c1.Val + jinwei
		jinwei = newVal / 10
		newNode := &ListNode{Val: newVal % 10, Next: nil}
		resLast.Next = newNode
		resLast = newNode
		c1 = c1.Next
	}

	for c2 != nil {
		newVal := c2.Val + jinwei
		jinwei = newVal / 10
		newNode := &ListNode{Val: newVal % 10, Next: nil}
		resLast.Next = newNode
		resLast = newNode
		c2 = c2.Next
	}

	if jinwei != 0 {
		newNode := &ListNode{Val: jinwei, Next: nil}
		resLast.Next = newNode
		resLast = newNode
	}

	return reverse11(resHead.Next)
}

func reverse11(h *ListNode) *ListNode {
	if h == nil || h.Next == nil {
		return h
	}
	var cur *ListNode
	next := h

	// nil -> 1 -> 2
	// nil <- 1 <- 2
	for next != nil {
		next.Next, cur, next = cur, next, next.Next
	}

	return cur
}
