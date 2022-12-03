package tree

import (
	"container/list"
)

/**
BM23 二叉树的前序遍历
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

数据范围：二叉树的节点数量满足 1 \le n \le 100 \1≤n≤100  ，二叉树节点的值满足 1 \le val \le 100 \1≤val≤100  ，树的各节点的值各不相同
示例 1：

}

*/

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	// return preorder1(root, res)
	// return preorder2(root, res)
	return preorder3(root, res)
}

func preorder1(root *TreeNode, nums []int) []int {
	if root != nil {
		nums = append(nums, root.Val)
		nums = preorder1(root.Left, nums)
		nums = preorder1(root.Right, nums)
	}
	return nums
}

func preorder2(root *TreeNode, nums []int) []int {
	if root == nil {
		return nums
	}

	stack := list.New()
	stack.PushBack(root)
	p := root
	for stack.Len() != 0 {
		p = stack.Remove(stack.Back()).(*TreeNode)
		nums = append(nums, p.Val)
		if p.Right != nil {
			stack.PushBack(p.Right)
		}
		if p.Left != nil {
			stack.PushBack(p.Left)
		}
	}

	return nums
}

func preorder3(root *TreeNode, nums []int) []int {
	if root == nil {
		return nums
	}

	stack := list.New()
	p := root
	for stack.Len() != 0 || p != nil {
		for p != nil {
			nums = append(nums, p.Val)
			stack.PushBack(p)
			p = p.Left
		}
		p = stack.Remove(stack.Back()).(*TreeNode)
		p = p.Right
	}

	return nums
}
