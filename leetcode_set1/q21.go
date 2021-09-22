package code_set1

/**
21. 合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。



示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var (
        res = &ListNode{
            Val:  0,
            Next: nil,
        }
        t1 = l1
        t2 = l2
        st = res
    )

    for t1 != nil && t2 != nil {
        st.Next = &ListNode{Val: 0, Next: nil}
        st = st.Next
        if t1.Val > t2.Val {
            st.Val = t2.Val
            t2 = t2.Next
        } else {
            st.Val = t1.Val
            t1 = t1.Next
        }
    }
    for t1 != nil {
        st.Next = &ListNode{Val: 0, Next: nil}
        st = st.Next
        st.Val = t1.Val
        t1 = t1.Next
    }
    for t2 != nil {
        st.Next = &ListNode{Val: 0, Next: nil}
        st = st.Next
        st.Val = t2.Val
        t2 = t2.Next
    }

    return res.Next

}*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    var head *ListNode
    if l1.Val < l2.Val {
        head = l1
        l1 = l1.Next
    } else {
        head = l2
        l2 = l2.Next
    }
    ptr := head
    for {
        if l1 == nil {
            ptr.Next = l2
            return head
        }
        if l2 == nil {
            ptr.Next = l1
            return head
        }
        if l1.Val < l2.Val {
            ptr.Next = l1
            l1 = l1.Next
        } else {
            ptr.Next = l2
            l2 = l2.Next
        }
        ptr = ptr.Next
    }
    return head
}
