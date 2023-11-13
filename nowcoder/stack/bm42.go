package stack

/**
BM42 用两个栈实现队列
题目
题解(371)
讨论(2k)
排行
面经new
简单  通过率：41.65%  时间限制：1秒  空间限制：256M
知识点
栈
描述
用两个栈来实现一个队列，使用n个元素来完成 n 次在队列尾部插入整数(push)和n次在队列头部删除整数(pop)的功能。 队列中的元素为int类型。保证操作合法，即保证pop操作时队列内已有元素。

数据范围： n\le1000n≤1000
要求：存储n个元素的空间复杂度为 O(n)O(n) ，插入与删除的时间复杂度都是 O(1)O(1)
*/

var stack421 []int // push
var stack422 []int // pop

func Push(node int) {
	stack421 = append(stack421, node)
}

func Pop() int {
	if len(stack422) == 0 {
		// assert len(stack(421) > 0
		for i := len(stack421) - 1; i >= 0; i-- {
			stack422 = append(stack422, stack421[i])
		}
		stack421 = []int{}
	}
	l := len(stack422)
	res := stack422[l-1]
	stack422 = stack422[0 : l-1]
	return res
}
