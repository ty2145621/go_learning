package stack

import (
	"container/list"
)

/**

BM43 包含min函数的栈
题目
题解(206)
讨论(1k)
排行
面经new
简单  通过率：35.46%  时间限制：1秒  空间限制：256M
知识点
栈
描述
定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的 min 函数，输入操作时保证 pop、top 和 min 函数操作时，栈中一定有元素。

此栈包含的方法有：
push(value):将value压入栈中
pop():弹出栈顶元素
top():获取栈顶元素
min():获取栈中最小元素

数据范围：操作数量满足 0 \le n \le 300 \0≤n≤300  ，输入的元素满足 |val| \le 10000 \∣val∣≤10000
进阶：栈的各个操作的时间复杂度是 O(1)\O(1)  ，空间复杂度是 O(n)\O(n)

示例:
输入:    ["PSH-1","PSH2","MIN","TOP","POP","PSH1","TOP","MIN"]
输出:    -1,2,1,-1
解析:
"PSH-1"表示将-1压入栈中，栈中元素为-1
"PSH2"表示将2压入栈中，栈中元素为2，-1
“MIN”表示获取此时栈中最小元素==>返回-1
"TOP"表示获取栈顶元素==>返回2
"POP"表示弹出栈顶元素，弹出2，栈中元素为-1
"PSH1"表示将1压入栈中，栈中元素为1，-1
"TOP"表示获取栈顶元素==>返回1
“MIN”表示获取此时栈中最小元素==>返回-1
*/

var stack43 = list.New()
var stack43Min = list.New()

func Push43(node int) {
	stack43.PushBack(node)
	if stack43Min.Len() == 0 || stack43Min.Back().Value.(int) > node {
		stack43Min.PushBack(node)
	} else {
		stack43Min.PushBack(stack43Min.Back().Value.(int))
	}
}
func Pop43() {
	stack43Min.Remove(stack43Min.Back())
	stack43.Remove(stack43.Back())
}
func Top() int {
	return stack43.Back().Value.(int)
}
func Min() int {
	return stack43Min.Back().Value.(int)
}
