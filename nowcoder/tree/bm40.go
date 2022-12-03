package tree

/**
BM40 重建二叉树
题目
题解(339)
讨论(2k)
排行
面经new
中等  通过率：27.77%  时间限制：1秒  空间限制：256M
知识点
树
dfs
数组
描述
给定节点数为 n 的二叉树的前序遍历和中序遍历结果，请重建出该二叉树并返回它的头结点。
例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建出如下图所示。


提示:
1.vin.length == pre.length
2.pre 和 vin 均无重复元素
3.vin出现的元素均出现在 pre里
4.只需要返回根结点，系统会自动输出整颗树做答案对比
数据范围：n \le 2000n≤2000，节点的值 -10000 \le val \le 10000−10000≤val≤10000
要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
*/

// 这题没写非递归的写法
func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	// write code here
	return reConstructBinaryTree1(pre, vin)
}

func reConstructBinaryTree1(pre []int, vin []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	p := &TreeNode{Val: pre[0]}
	if len(pre) == 1 {
		return p
	}

	for i := 0; i < len(vin); i++ {
		if pre[0] == vin[i] {
			p.Left = reConstructBinaryTree1(pre[1:i+1], vin[0:i])
			if i != len(vin)-1 {
				p.Right = reConstructBinaryTree1(pre[i+1:], vin[i+1:])
			}
		}
	}
	return p
}
