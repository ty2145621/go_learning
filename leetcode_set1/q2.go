package leetcode_set1

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		jinwei           = 0
		node   *ListNode = new(ListNode)
		head             = node
		l1Head           = l1
		l2Head           = l2
		tNode  *ListNode = nil
	)
	for l1Head != nil && l2Head != nil {
		node.Next = new(ListNode)
		node = node.Next
		node.Val = jinwei + l1Head.Val + l2Head.Val
		if node.Val >= 10 {
			node.Val = node.Val - 10
			jinwei = 1
		} else {
			jinwei = 0
		}
		l1Head = l1Head.Next
		l2Head = l2Head.Next
	}

	if l1Head != nil {
		tNode = l1Head
	}
	if l2Head != nil {
		tNode = l2Head
	}

	for tNode != nil {
		node.Next = new(ListNode)
		node = node.Next
		node.Val = jinwei + tNode.Val
		if node.Val >= 10 {
			node.Val = node.Val - 10
			jinwei = 1
		} else {
			jinwei = 0
		}
		tNode = tNode.Next
	}

	if jinwei != 0 {
		node.Next = new(ListNode)
		node = node.Next
		node.Val = jinwei
	}

	return head.Next
}
