package tree

import (
	"container/list"
)

/**
BM33 二叉树的镜像

操作给定的二叉树，将其变换为源二叉树的镜像。
数据范围：二叉树的节点数 0 \le n \le 10000≤n≤1000 ， 二叉树每个节点的值 0\le val \le 10000≤val≤1000
要求： 空间复杂度 O(n)O(n) 。本题也有原地操作，即空间复杂度 O(1)O(1) 的解法，时间复杂度 O(n)O(n)
*/

func Mirror(pRoot *TreeNode) *TreeNode {
	// return mirror1(pRoot)
	return mirror2(pRoot)
}

func mirror1(h *TreeNode) *TreeNode {
	if h == nil {
		return h
	}
	h.Left, h.Right = h.Right, h.Left
	mirror1(h.Left)
	mirror1(h.Right)

	return h
}

func mirror2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() != 0 {
		h := stack.Remove(stack.Front()).(*TreeNode)
		if h == nil {
			continue
		}
		h.Left, h.Right = h.Right, h.Left
		stack.PushBack(h.Left)
		stack.PushBack(h.Right)
	}

	return root
}
