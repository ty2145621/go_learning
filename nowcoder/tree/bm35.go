package tree

import (
	"container/list"
)

/**
BM35 判断是不是完全二叉树

给定一个二叉树，确定他是否是一个完全二叉树。

完全二叉树的定义：若二叉树的深度为 h，除第 h 层外，其它各层的结点数都达到最大个数，第 h 层所有的叶子结点都连续集中在最左边，这就是完全二叉树。（第 h 层可能包含 [1~2h] 个节点）
*/

func isCompleteTree(root *TreeNode) bool {
	// write code here
	stack := list.New()
	stack.PushBack(root)
	findLast := false
	for stack.Len() > 0 {
		size := stack.Len()
		for i := 0; i < size; i++ {
			p := stack.Remove(stack.Front()).(*TreeNode)
			if p == nil {
				findLast = true
				continue
			}
			if findLast && p != nil {
				return false
			}
			stack.PushBack(p.Left)
			stack.PushBack(p.Right)
		}
	}
	return true
}
