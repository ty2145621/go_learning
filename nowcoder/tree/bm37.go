package tree

/*
*
BM37 二叉搜索树的最近公共祖先
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
1.对于该题的最近的公共祖先定义:对于有根树T的两个节点p、q，最近公共祖先LCA(T,p,q)表示一个节点x，满足x是p和q的祖先且x的深度尽可能大。在这里，一个节点也可以是它自己的祖先.
2.二叉搜索树是若它的左子树不空，则左子树上所有节点的值均小于它的根节点的值； 若它的右子树不空，则右子树上所有节点的值均大于它的根节点的值
3.所有节点的值都是唯一的。
4.p、q 为不同节点且均存在于给定的二叉搜索树中。
数据范围:
3<=节点总数<=10000
0<=节点值<=10000
*/

func lowestCommonAncestor(root *TreeNode, p int, q int) int {
	// write code here
	resp := findPath(root, p)
	resq := findPath(root, q)

	res := resp[0]
	for i := 0; i < len(resp) && i < len(resq); i++ {
		if resp[i] == resq[i] {
			res = resp[i]
		}
	}
	return res
}

func findPath(root *TreeNode, s int) []int {
	res := make([]int, 0)

	p := root
	for p != nil {
		res = append(res, p.Val)
		if p.Val < s {
			p = p.Right
		} else if p.Val == s {
			return res
		} else if p.Val > s {
			p = p.Left
		}
	}
	return res
}
