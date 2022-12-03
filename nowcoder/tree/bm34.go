package tree

import (
	"container/list"
	"math"
)

/**
BM34 判断是不是二叉搜索树
给定一个二叉树根节点，请你判断这棵树是不是二叉搜索树。

二叉搜索树满足每个节点的左子树上的所有节点均小于当前节点且右子树上的所有节点均大于当前节点。

*/

func isValidBST(root *TreeNode) bool {
	// write code here
	return isValidBST1(root, math.MinInt32, math.MaxInt32)
}

func isValidBST1(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && (root.Left.Val > root.Val || root.Left.Val < min) {
		return false
	}
	if root.Right != nil && (root.Right.Val < root.Val || root.Right.Val > max) {
		return false
	}
	return isValidBST1(root.Left, min, root.Val) && isValidBST1(root.Right, root.Val, max)
}

// 中序遍历
func isValidBST2(root *TreeNode) bool {
	stack := list.New()
	p := root
	pre := math.MinInt32
	for p != nil || stack.Len() > 0 {
		for p != nil {
			stack.PushBack(p)
			p = p.Left
		}
		p = stack.Remove(stack.Back()).(*TreeNode)
		if p.Val < pre {
			return false
		}
		pre = p.Val
		p = p.Right
	}
	return true
}
