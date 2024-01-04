package code

import (
	"container/list"
	"strconv"
)

// k个一组翻转链表

type listNode struct {
	val  int
	next *listNode
}

func build(arr []int) *listNode {
	head := &listNode{}
	cur := head
	for _, v := range arr {
		cur.next = &listNode{
			val:  v,
			next: nil,
		}
		cur = cur.next
	}
	return head.next
}

func (l *listNode) String() string {
	cur := l
	res := ""
	for cur != nil {
		res += strconv.Itoa(cur.val) + ","
		cur = cur.next
	}
	return res
}

func reverseKGroup(head *listNode, k int) *listNode {
	newHead := &listNode{next: head}
	pre := newHead

	for pre != nil && pre.next != nil {
		if !canReverse(pre.next, k) {
			return newHead.next
		}
		pre.next, pre = reverseKGroupPartition(pre.next, k), pre.next
	}
	return newHead.next
}

// reverseKGroupWithLast 最后不足K个也要翻转
func reverseKGroupWithLast(head *listNode, k int) *listNode {
	newHead := &listNode{next: head}
	pre := newHead

	for pre != nil && pre.next != nil {
		pre.next, pre = reverseKGroupPartition(pre.next, k), pre.next
	}
	return newHead.next
}

// canReverse true表示可以翻转
func canReverse(head *listNode, k int) bool {
	cur := head
	for ; k > 0; k-- {
		if cur == nil {
			return false
		}
		cur = cur.next
	}
	return true
}

func reverseKGroupPartition(head *listNode, k int) *listNode {
	var pre *listNode
	cur := head
	for ; k > 0; k-- { // 不需要判断cur，必定非nil
		if cur == nil {
			break
		}
		pre, cur, cur.next = cur, cur.next, pre
	}

	head.next = cur
	return pre
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	// return preorderTraversal1(root)
	return preorderTraversal2(root)
}

// 非递归
func preorderTraversal2(root *TreeNode) []int {
	arr := make([]int, 0, 256)
	if root == nil {
		return arr
	}
	cur := root
	stack := list.New() // *TreeNode
	stack.PushBack(cur.Right)
	for stack.Len() > 0 {
		for cur != nil {
			arr = append(arr, cur.Val)
			stack.PushBack(cur.Right)
			cur = cur.Left
		}
		cur = stack.Remove(stack.Back()).(*TreeNode)
	}
	return arr
}

// 递归
func preorderTraversal1(root *TreeNode) []int {
	cur := root
	arr := make([]int, 0, 256)
	arr = preorderTraversal1P(cur, arr)
	return arr
}

func preorderTraversal1P(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}
	arr = append(arr, root.Val)
	arr = preorderTraversal1P(root.Left, arr)
	arr = preorderTraversal1P(root.Right, arr)
	return arr
}

func inorderTraversal(root *TreeNode) []int {
	return inorderTraversal1(root)
	// return inorderTraversal2(root)
}

func inorderTraversal1(root *TreeNode) []int {
	arr := make([]int, 0, 256)
	return inorderTraversal1P(root, arr)
}

func inorderTraversal1P(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}
	arr = inorderTraversal1P(root.Left, arr)
	arr = append(arr, root.Val)
	arr = inorderTraversal1P(root.Right, arr)
	return arr
}

func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	cur := root
	arr := make([]int, 0, 256)
	stack := list.New()
	for stack.Len() > 0 || cur != nil {
		for cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		}
		cur = stack.Remove(stack.Back()).(*TreeNode)
		arr = append(arr, cur.Val)
		cur = cur.Right
	}
	return arr
}

func postorderTraversal(root *TreeNode) []int {
	return postorderTraversal1(root)
}

func postorderTraversal1(root *TreeNode) []int {
	arr := make([]int, 0, 256)
	return postorderTraversal1P(root, arr)
}

func postorderTraversal1P(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}
	arr = postorderTraversal1P(root.Left, arr)
	arr = postorderTraversal1P(root.Right, arr)
	arr = append(arr, root.Val)
	return arr
}

func postorderTraversal2(root *TreeNode) []int {
	arr := make([]int, 0, 256)
	cur := root
	stack := list.New()
	viewMap := make(map[*TreeNode]bool)

	for stack.Len() > 0 || cur != nil {
		for cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		}
		cur = stack.Back().Value.(*TreeNode)
		for viewMap[cur] {
			arr = append(arr, cur.Val)
			stack.Remove(stack.Back())
			if stack.Len() == 0 {
				return arr
			}
			cur = stack.Back().Value.(*TreeNode)
		}
		viewMap[cur] = true
		cur = cur.Right
	}
	return arr
}

func levelOrder(root *TreeNode) [][]int {
	arrs := make([][]int, 0, 256)
	stack := list.New()
	if root == nil {
		return arrs
	}
	stack.PushBack(root)
	for stack.Len() > 0 {
		levelLen := stack.Len()
		arr := make([]int, 0)
		for i := 0; i < levelLen; i++ {
			cur := stack.Remove(stack.Front()).(*TreeNode)
			if cur == nil {
				continue
			}
			arr = append(arr, cur.Val)
			if cur.Left != nil {
				stack.PushBack(cur.Left)
			}
			if cur.Right != nil {
				stack.PushBack(cur.Right)
			}
		}
		arrs = append(arrs, arr)
	}
	return arrs
}

func lowestCommonAncestor(root *TreeNode, p int, q int) int {
	pl := findPath(root, p)
	ql := findPath(root, q)

	last := -1
	for pl.Len() > 0 && ql.Len() > 0 {
		pLast := pl.Remove(pl.Front()).(*TreeNode)
		qLast := ql.Remove(ql.Front()).(*TreeNode)
		if pLast.Val == qLast.Val {
			last = pLast.Val
		} else {
			break
		}
	}
	return last
}

func findPath(root *TreeNode, n int) *list.List {
	s := list.New()
	path := list.New()
	cur := root

	for s.Len() > 0 || cur != nil {
		for cur != nil {
			s.PushBack(cur)
			path.PushBack(cur)
			if cur.Val == n {
				return path
			}
			cur = cur.Left
		}
		cur = s.Remove(s.Back()).(*TreeNode)
		for path.Back().Value.(*TreeNode) != cur {
			path.Remove(path.Back())
		}
		cur = cur.Right
	}
	return path
}
