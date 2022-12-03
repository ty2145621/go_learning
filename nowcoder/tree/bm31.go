package tree

import (
	"container/list"
)

/**
BM31 对称的二叉树
描述
给定一棵二叉树，判断其是否是自身的镜像（即：是否对称）
例如：                                 下面这棵二叉树是对称的

下面这棵二叉树不对称。

数据范围：节点数满足 0 \le n \le 10000≤n≤1000，节点上的值满足 |val| \le 1000∣val∣≤1000
要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
备注：
你可以用递归和迭代两种方法解决这个问题
*/

func isSymmetrical(pRoot *TreeNode) bool {
	// return isSymmetrical1(pRoot)
	return isSymmetrical2(pRoot)
}

func isSymmetrical1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}

func isSymmetrical2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := list.New()
	l.PushBack(root.Left)
	r := list.New()
	r.PushBack(root.Right)
	for l.Len() != 0 || r.Len() != 0 {
		if l.Len() != r.Len() {
			return false
		}
		size := l.Len()
		for ; size > 0; size-- {
			lp := l.Remove(l.Front()).(*TreeNode)
			rp := r.Remove(r.Front()).(*TreeNode)
			if lp == nil && rp == nil {
				continue
			}
			if lp == nil || rp == nil {
				return false
			}
			if lp.Val != rp.Val {
				return false
			}
			l.PushBack(lp.Left)
			l.PushBack(lp.Right)
			r.PushBack(rp.Right)
			r.PushBack(rp.Left)
		}
	}
	return true
}
