package tree

import (
	"container/list"
)

/**
BM25 二叉树的后序遍历
*/

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	// return postorder1(root, res)
	// return postorder2(root, res)
	return postorder3(root, res)
}

func postorder1(root *TreeNode, nums []int) []int {
	if root == nil {
		return nums
	}
	nums = postorder1(root.Left, nums)
	nums = postorder1(root.Right, nums)
	nums = append(nums, root.Val)
	return nums
}

// postorder2 需要存储已经遍历过右子树的父节点
func postorder2(root *TreeNode, nums []int) []int {
	if root == nil {
		return nums
	}

	m := make(map[*TreeNode]struct{})
	stack := list.New()
	p := root
	for stack.Len() != 0 || p != nil {
		for p != nil {
			stack.PushBack(p)
			p = p.Left
		}
		top := stack.Back().Value.(*TreeNode)
		if _, ok := m[top]; ok {
			nums = append(nums, top.Val)
			stack.Remove(stack.Back())
			p = nil
		} else {
			p = top.Right
			m[top] = struct{}{}
		}
	}

	return nums
}

// postorder3 相比2，实际上只需要存储本次处理的子树的root节点就好
func postorder3(root *TreeNode, nums []int) []int {
	if root == nil {
		return nums
	}

	stack := list.New()
	last := root
	p := root
	for stack.Len() != 0 || p != nil {
		for p != nil {
			stack.PushBack(p)
			p = p.Left
		}
		top := stack.Back().Value.(*TreeNode)
		if top.Right == nil || top.Right == last {
			nums = append(nums, top.Val)
			stack.Remove(stack.Back())
			p = nil
		} else {
			p = top.Right
		}
		last = top
	}

	return nums
}
