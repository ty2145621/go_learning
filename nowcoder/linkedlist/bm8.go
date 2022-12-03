package linkedlist

/**
BM8 链表中倒数最后k个结点
输入一个长度为 n 的链表，设链表中的元素的值为 ai ，返回该链表中倒数第k个节点。
如果该链表长度小于k，请返回一个长度为 0 的链表。


数据范围：0 \leq n \leq 10^50≤n≤10
5
 ，0 \leq a_i \leq 10^90≤a
i
​
 ≤10
9
 ，0 \leq k \leq 10^90≤k≤10
9

要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
进阶：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)

例如输入{1,2,3,4,5},2时，对应的链表结构如下图所示：

其中蓝色部分为该链表的最后2个结点，所以返回倒数第2个结点（也即结点值为4的结点）即可，系统会打印后面所有的节点来比较。
示例1
输入：
{1,2,3,4,5},2
复制
返回值：
{4,5}
复制
说明：
返回倒数第2个节点4，系统会打印后面所有的节点来比较。
示例2
输入：
{2},8
复制
返回值：
{}

*/

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	cur := pHead
	fast := pHead
	for k > 0 && fast != nil {
		k--
		fast = fast.Next
	}

	if k > 0 {
		return nil
	}

	for fast != nil {
		cur = cur.Next
		fast = fast.Next
	}
	return cur
}
