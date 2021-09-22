package leetcode_set1

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
    if k == 1 {
        return head
    }
    tmpHead := &ListNode{}
    tmpHead.Next = head

    var tHead = tmpHead
    var tEnd = getAfterK(tHead.Next, k)

    for tEnd != nil {
        h := tHead.Next
        e := tEnd
        reverseK(h, e)
        tHead.Next = e
        tHead = h
        tEnd = getAfterK(tHead.Next, k)
    }

    return tmpHead.Next
}

func getAfterK(begin *ListNode, k int) *ListNode {
    var t = begin
    for ; k > 1; k-- {
        if t != nil {
            t = t.Next
        } else {
            break
        }
    }
    return t
}

// 0->1->2->3->4->5->6->7->8
func reverseK(head *ListNode, end *ListNode) {
    tmpHead := &ListNode{}
    tmpHead.Next = head

    // assert 链表长度>=k
    // assert end != nil
    ePtr := end.Next
    nPtr := tmpHead.Next
    cPtr := nPtr.Next

    for {
        nPtr.Next = ePtr
        ePtr = nPtr
        nPtr = cPtr
        cPtr = cPtr.Next

        if nPtr == end {
            nPtr.Next = ePtr
            tmpHead.Next = nPtr
            break
        }
    }
    return
}
