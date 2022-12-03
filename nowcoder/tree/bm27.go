package tree

import (
	"container/list"
)

/**
BM27 按之字形顺序打印二叉树
题目
题解(301)
讨论(1k)
排行
面经new
中等  通过率：29.60%  时间限制：1秒  空间限制：256M
知识点
栈
树
队列
描述
给定一个二叉树，返回该二叉树的之字形层序遍历，（第一层从左向右，下一层从右向左，一直这样交替）

数据范围：0 \le n \le 15000≤n≤1500,树上每个节点的val满足 |val| <= 1500∣val∣<=1500
要求：空间复杂度：O(n)O(n)，时间复杂度：O(n)O(n)
例如：
给定的二叉树是{1,2,3,#,#,4,5}
*/

func Print(pRoot *TreeNode) [][]int {
	res := make([][]int, 0)
	if pRoot == nil {
		return res
	}

	p := pRoot
	l := list.New()
	l.PushBack(p)

	level := 0
	for l.Len() != 0 {
		size := l.Len()
		tmp := make([]int, size)
		for ; size > 0; size-- {
			p = l.Remove(l.Front()).(*TreeNode)
			if level%2 == 0 {
				tmp[len(tmp)-size] = p.Val
			} else {
				tmp[size-1] = p.Val
			}
			if p.Left != nil {
				l.PushBack(p.Left)
			}
			if p.Right != nil {
				l.PushBack(p.Right)
			}
		}
		res = append(res, tmp)

		level++
	}
	return res
}
