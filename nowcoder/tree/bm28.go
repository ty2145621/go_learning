package tree

import (
	"container/list"
)

/**
BM28 二叉树的最大深度

描述
求给定二叉树的最大深度，
深度是指树的根节点到任一叶子节点路径上节点的数量。
最大深度是所有叶子节点的深度的最大值。
（注：叶子节点是指没有子节点的节点。）


数据范围：0 \le n \le 1000000≤n≤100000，树上每个节点的val满足 |val| \le 100∣val∣≤100
要求： 空间复杂度 O(1)O(1),时间复杂度 O(n)O(n)
*/

func maxDepth(root *TreeNode) int {
	// write code here
	// return maxDepth1(root)
	return maxDepth2(root)
}

func maxDepth1(root *TreeNode) int {
	if root != nil {
		return getMax(maxDepth1(root.Left), maxDepth1(root.Right)) + 1
	}
	return 0
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := list.New()
	l.PushBack(root)
	var p *TreeNode
	level := 0
	for l.Len() != 0 {
		size := l.Len()
		for ; size > 0; size-- {
			p = l.Remove(l.Front()).(*TreeNode)
			if p.Left != nil {
				l.PushBack(p.Left)
			}
			if p.Right != nil {
				l.PushBack(p.Right)
			}
		}
		level++
	}
	return level
}
