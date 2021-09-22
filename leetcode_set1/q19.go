package code_set1

/**
19. 删除链表的倒数第N个节点
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // 多加一个头节点方便处理head
    lhead := &ListNode{
        Val:  0,
        Next: head,
    }
    var (
        start = lhead
        end   = lhead
    )

    for ; n > 0 && end != nil; n-- {
        end = end.Next
    }

    for end != nil && end.Next != nil {
        end = end.Next
        start = start.Next
    }

    // delete
    start.Next = start.Next.Next

    return lhead.Next
}
