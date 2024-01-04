package interview

import (
	"container/list"
)

// 二叉树的后序遍历
type listNode struct {
	val   int
	left  *listNode
	right *listNode
}

func postOrder(root *listNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	cur := root
	stack := list.New()
	last := root
	for stack.Len() != 0 || cur != nil {
		for cur != nil {
			stack.PushBack(cur)
			cur = cur.left
		}
		// cur = stack.Remove(stack.Back()).(*listNode)
		top := stack.Back().Value.(*listNode)
		if top.right == nil || top.right == last {
			res = append(res, top.val)
			stack.Remove(stack.Back())
			cur = nil
		} else {
			cur = top.right
		}
		last = top
	}

	return res
}

func postOrder2(root *listNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	cur := root
	stack := list.New()
	visit := make(map[*listNode]bool) // 是否访问过右节点
	for stack.Len() != 0 || cur != nil {
		for cur != nil {
			stack.PushBack(cur)
			cur = cur.left
		}
		top := stack.Back().Value.(*listNode)
		if visit[top] {
			res = append(res, top.val)
			stack.Remove(stack.Back())
			cur = nil
		} else {
			visit[cur] = true
			cur = top.right
		}
	}
	return res
}
