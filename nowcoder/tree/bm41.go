package tree

import (
	"container/list"
)

/*
*
BM41 输出二叉树的右视图
题目
题解(122)
讨论(150)
排行
面经new
中等  通过率：59.32%  时间限制：1秒  空间限制：256M
知识点
树
描述
请根据二叉树的前序遍历，中序遍历恢复二叉树，并打印出二叉树的右视图

数据范围： 0 \le n \le 100000≤n≤10000
要求： 空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)

如输入[1,2,4,5,3],[4,2,5,1,3]时，通过前序遍历的结果[1,2,4,5,3]和中序遍历的结果[4,2,5,1,3]可重建出以下二叉树：

所以对应的输出为[1,3,5]。
示例1
输入：
[1,2,4,5,3],[4,2,5,1,3]
复制
返回值：
[1,3,5]
*/

// solve 先构造数，接着层序遍历，获取层序的最后一个节点
func solve(pre []int, inorder []int) []int {
	res := make([]int, 0)
	root := reConstructBinaryTree1(pre, inorder)
	s := list.New()
	s.PushBack(root)
	for s.Len() > 0 {
		size := s.Len()
		for ; size > 0; size-- {
			p := s.Remove(s.Front()).(*TreeNode)
			if p.Left != nil {
				s.PushBack(p.Left)
			}
			if p.Right != nil {
				s.PushBack(p.Right)
			}
			if size == 1 {
				res = append(res, p.Val)
			}
		}
	}
	return res
}
