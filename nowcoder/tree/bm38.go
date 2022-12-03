package tree

import (
	"container/list"
)

/**
BM38 在二叉树中找到两个节点的最近公共祖先
给定一棵二叉树(保证非空)以及这棵树上的两个节点对应的val值 o1 和 o2，请找到 o1 和 o2 的最近公共祖先节点。

数据范围：树上节点数满足 1 \le n \le 10^5 \1≤n≤10
5
   , 节点值val满足区间 [0,n)
要求：时间复杂度 O(n)O(n)

注：本题保证二叉树中每个节点的val值均不相同。

如当输入{3,5,1,6,2,0,8,#,#,7,4},5,1时，二叉树{3,5,1,6,2,0,8,#,#,7,4}如下图所示：
*/

func lowestCommonAncestor38(root *TreeNode, o1 int, o2 int) int {
	// write code here
	// return lowestCommonAncestor1(root, o1, o2)
	a1 := findPath38(root, o1)
	a2 := findPath38(root, o2)

	res := a1.Front().Value.(int)
	for a1.Len() > 0 && a2.Len() > 0 {
		t1 := a1.Remove(a1.Front()).(int)
		t2 := a2.Remove(a2.Front()).(int)
		if t1 == t2 {
			res = t1
		} else {
			break
		}
	}
	return res
}

func lowestCommonAncestor1(root *TreeNode, o1, o2 int) int {
	if root == nil {
		return -1
	}
	if root.Val == o1 || root.Val == o2 {
		return root.Val
	}

	left := lowestCommonAncestor1(root.Left, o1, o2)
	right := lowestCommonAncestor1(root.Right, o1, o2)
	if left != -1 && right != -1 {
		return root.Val
	}
	if left == -1 {
		return right
	}
	return left
}

func findPath38(root *TreeNode, target int) *list.List {
	s := list.New()
	path := list.New()

	p := root
	for s.Len() > 0 || p != nil {
		for p != nil {
			s.PushBack(p)
			path.PushBack(p.Val)
			p = p.Left
		}
		p = s.Remove(s.Back()).(*TreeNode)
		for p.Val != path.Back().Value.(int) {
			path.Remove(path.Back())
		}
		if p.Val == target {
			return path
		}
		p = p.Right
	}
	return path
}
