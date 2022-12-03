package tree

import (
	"container/list"
	"strconv"
	"strings"
)

/**
BM39 序列化二叉树
题目
题解(155)
讨论(577)
排行
面经new
较难  通过率：25.31%  时间限制：1秒  空间限制：256M
知识点
队列
树
描述
请实现两个函数，分别用来序列化和反序列化二叉树，不对序列化之后的字符串进行约束，但要求能够根据序列化之后的字符串重新构造出一棵与原二叉树相同的树。

二叉树的序列化(Serialize)是指：把一棵二叉树按照某种遍历方式的结果以某种格式保存为字符串，从而使得内存中建立起来的二叉树可以持久保存。序列化可以基于先序、中序、后序、层序的二叉树等遍历方式来进行修改，序列化的结果是一个字符串，序列化时通过 某种符号表示空节点（#）

二叉树的反序列化(Deserialize)是指：根据某种遍历顺序得到的序列化字符串结果str，重构二叉树。

例如，可以根据层序遍历的方案序列化，如下图:

层序序列化(即用函数Serialize转化)如上的二叉树转为"{1,2,3,#,#,6,7}"，再能够调用反序列化(Deserialize)将"{1,2,3,#,#,6,7}"构造成如上的二叉树。

当然你也可以根据满二叉树结点位置的标号规律来序列化，还可以根据先序遍历和中序遍历的结果来序列化。不对序列化之后的字符串进行约束，所以欢迎各种奇思妙想。

数据范围：节点数 n \le 100n≤100，树上每个节点的值满足 0 \le val \le 1500≤val≤150
要求：序列化和反序列化都是空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
*/

func Serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	s := list.New()
	s.PushBack(root)
	res := make([]string, 0)
	for s.Len() > 0 {
		p := s.Remove(s.Front()).(*TreeNode)
		if p == nil {
			res = append(res, "-1")
		} else {
			res = append(res, strconv.Itoa(p.Val))
			s.PushBack(p.Left)
			s.PushBack(p.Right)
		}
	}
	return strings.Join(res, ",")
}
func Deserialize(s string) *TreeNode {
	arr := strings.Split(s, ",")
	if len(arr) == 0 {
		return nil
	}
	arrInt := make([]int, 0)
	for _, v := range arr {
		i, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		arrInt = append(arrInt, i)
	}
	if len(arrInt) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: arrInt[0],
	}
	stack := list.New()
	stack.PushBack(root)
	for i := 1; i < len(arrInt); i++ {
		p := stack.Remove(stack.Front()).(*TreeNode)
		if arrInt[i] != -1 {
			p.Left = &TreeNode{
				Val: arrInt[i],
			}
			stack.PushBack(p.Left)
		}
		i++
		if i < len(arrInt) && arrInt[i] != -1 {
			p.Right = &TreeNode{
				Val: arrInt[i],
			}
			stack.PushBack(p.Right)
		}
	}
	return root
}
