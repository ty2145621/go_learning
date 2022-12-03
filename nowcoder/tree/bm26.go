package tree

import (
	"container/list"
)

/**
BM26 求二叉树的层序遍历
*/

var nums26 [][]int

func levelOrder(root *TreeNode) [][]int {
	nums26 = make([][]int, 0)

	// levelOrder1(root, 0)
	levelOrder2(root, 0)
	return nums26
}

func levelOrder1(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if len(nums26) == level {
		nums26 = append(nums26, make([]int, 0))
	}

	nums26[level] = append(nums26[level], root.Val)
	levelOrder1(root.Left, level+1)
	levelOrder1(root.Right, level+1)
	return
}

func levelOrder2(root *TreeNode, level int) [][]int {
	if root == nil {
		return nums26
	}
	l := list.New()
	l.PushBack(root)

	for l.Len() != 0 {
		size := l.Len()
		nums26 = append(nums26, make([]int, 0, size))
		for ; size > 0; size-- {
			p := l.Remove(l.Front()).(*TreeNode)
			nums26[level] = append(nums26[level], p.Val)
			if p.Left != nil {
				l.PushBack(p.Left)
			}
			if p.Right != nil {
				l.PushBack(p.Right)
			}
		}
		level++
	}
	return nums26
}
