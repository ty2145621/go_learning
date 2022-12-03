package tree

import (
	"math"
)

/**
BM36 判断是不是平衡二叉树

输入一棵节点数为 n 二叉树，判断该二叉树是否是平衡二叉树。
在这里，我们只需要考虑其平衡性，不需要考虑其是不是排序二叉树
平衡二叉树（Balanced Binary Tree），具有以下性质：它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树。
样例解释：
*/

func IsBalanced_Solution(pRoot *TreeNode) bool {
	// write code here
	ok, _ := IsBalanced_Solution1(pRoot)
	return ok
}

func IsBalanced_Solution1(root *TreeNode) (is bool, depth int) {
	if root == nil {
		return true, 0
	}
	isL, depL := IsBalanced_Solution1(root.Left)
	isR, depR := IsBalanced_Solution1(root.Right)
	if math.Abs(float64(depR-depL)) > 1 {
		return false, 0
	}
	return isL && isR, getMax(depL, depR) + 1
}
