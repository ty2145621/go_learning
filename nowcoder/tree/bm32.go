package tree

import (
	"container/list"
)

/**
BM32 合并二叉树

已知两颗二叉树，将它们合并成一颗二叉树。合并规则是：都存在的结点，就将结点值加起来，否则空的位置就由另一个树的结点来代替。例如：
两颗二叉树是:
*/

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	// return mergeTrees1(t1, t2)
	return mergeTrees1(t1, t2)
}

// mergeTrees1 t1 是最终的tree
func mergeTrees1(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val
	t1.Left = mergeTrees1(t1.Left, t2.Left)
	t1.Right = mergeTrees1(t1.Right, t2.Right)
	return t1
}

// mergeTrees2 t1 是最终的tree 非递归
func mergeTrees2(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	l1 := list.New()
	l2 := list.New()
	l1.PushBack(t1)
	l2.PushBack(t2)

	for l1.Len() != 0 && l2.Len() != 0 {
		p1 := l1.Remove(l1.Front()).(*TreeNode)
		p2 := l2.Remove(l2.Front()).(*TreeNode)
		p1.Val += p2.Val
		if p1.Left != nil && p2.Left != nil {
			l1.PushBack(p1.Left)
			l2.PushBack(p2.Left)
		} else {
			if p1.Left == nil {
				p1.Left, p2.Left = p2.Left, p1.Left
			} // 其余情况都是empty，也不需要再进行层序遍历了
		}

		if p1.Right != nil && p2.Right != nil {
			l1.PushBack(p1.Right)
			l2.PushBack(p2.Right)
		} else {
			if p1.Right == nil {
				p1.Right, p2.Right = p2.Right, p1.Right
			} // 其余情况都是empty，也不需要再进行层序遍历了
		}
	}
	return t1
}
