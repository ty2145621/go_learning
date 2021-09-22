package leetcode_set1

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    tmpHead := &ListNode{}
    tmpHead.Next = head

    if head == nil || head.Next == nil {
        return head
    }

    var cur = tmpHead
    var next = cur.Next

    for {
        cur.Next = next.Next
        next.Next = cur.Next.Next
        cur.Next.Next = next

        cur = next
        next = cur.Next

        if next == nil || next.Next == nil {
            break
        }
    }

    return tmpHead.Next
}