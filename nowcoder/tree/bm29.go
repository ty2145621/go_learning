package tree

import (
	"container/list"
	"fmt"
)

/**
BM29 二叉树中和为某一值的路径(一)
描述
给定一个二叉树root和一个值 sum ，判断是否有从根节点到叶子节点的节点值之和等于 sum 的路径。
1.该题路径定义为从树的根结点开始往下一直到叶子结点所经过的结点
2.叶子节点是指没有子节点的节点
3.路径只能从父节点到子节点，不能从子节点到父节点
4.总节点数目为n

例如：
给出如下的二叉树，\ sum=22 sum=22，

返回true，因为存在一条路径 5\to 4\to 11\to 25→4→11→2的节点值之和为 22

数据范围：
1.树上的节点数满足 0 \le n \le 100000≤n≤10000
2.每 个节点的值都满足 |val| \le 1000∣val∣≤1000
要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
进阶：空间复杂度 O(树的高度)O(树的高度)，时间复杂度 O(n)O(n)
*/

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	// write code here
	// return preorderPathSum1(root, sum)
	// return preorderPathSum2(root, sum)
	r, nums := preorderPathSum3(root, sum)
	fmt.Printf("%+v\n", nums)
	return r
}

func preorderPathSum1(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	target -= root.Val
	if target == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	return preorderPathSum1(root.Left, target) || preorderPathSum1(root.Right, target)
}

func preorderPathSum2(root *TreeNode, target int) bool {
	resList := list.New()
	sum := 0

	if root == nil {
		return false
	}

	stack := list.New()
	stack.PushBack(root)
	sum += root.Val
	resList.PushBack(sum)
	p := root
	for stack.Len() != 0 {
		p = stack.Remove(stack.Back()).(*TreeNode)
		sum = resList.Remove(resList.Back()).(int)

		if target == sum && p.Left == nil && p.Right == nil {
			return true
		}

		if p.Right != nil {
			stack.PushBack(p.Right)
			resList.PushBack(sum + p.Right.Val)
		}
		if p.Left != nil {
			stack.PushBack(p.Left)
			resList.PushBack(sum + p.Left.Val)
		}
	}
	return false
}

// preorderPathSum3 返回路径
func preorderPathSum3(root *TreeNode, target int) (bool, []int) {
	resList := list.New()
	path := list.New()
	sum := 0

	if root == nil {
		return false, []int{}
	}

	stack := list.New()
	p := root
	for stack.Len() != 0 || p != nil {
		for p != nil {
			stack.PushBack(p)
			sum += p.Val
			resList.PushBack(sum)
			path.PushBack(p)
			p = p.Left
		}
		p = stack.Remove(stack.Back()).(*TreeNode)
		for path.Back().Value.(*TreeNode) != p {
			path.Remove(path.Back())
		}
		sum = resList.Remove(resList.Back()).(int)

		if target == sum && p.Left == nil && p.Right == nil {
			return true, list2Slice(path)
		}

		p = p.Right
	}
	return false, []int{}
}

func list2Slice(resList *list.List) []int {
	res := make([]int, 0)
	for resList.Len() != 0 {
		t := resList.Remove(resList.Front()).(*TreeNode)
		res = append(res, t.Val)
	}
	return res
}
