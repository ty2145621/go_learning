package tree

import (
	"container/list"
)

/**
BM24 二叉树的中序遍历
*/

func inorderTraversal(root *TreeNode) []int {
	// write code here
	res := make([]int, 0)
	// return inorder1(root, res)
	return inorder2(root, res)
}

func inorder1(root *TreeNode, nums []int) []int {
	if root == nil {
		return nums
	}

	nums = inorder1(root.Left, nums)
	nums = append(nums, root.Val)
	nums = inorder1(root.Right, nums)
	return nums
}
func inorder2(root *TreeNode, nums []int) []int {
	if root == nil {
		return nums
	}

	stack := list.New()
	p := root
	for stack.Len() != 0 || p != nil {
		for p != nil {
			stack.PushBack(p)
			p = p.Left
		}
		p = stack.Remove(stack.Back()).(*TreeNode)
		nums = append(nums, p.Val)
		p = p.Right
	}

	return nums
}
