package code

import (
	"strconv"
	"strings"
)

/**
给定数字n，遍历1-n，并且要求按以下规则打印字符串：

如果是3的倍数，打印"a"
如果是5的倍数，打印"b"
如果是3和5的倍数，打印"ab"
其他情况打印自己

比如给定数字5，那么打印结果是：

1
2
a
4
b

要求：
设计一个package/lib并暴露接口，实现需要灵活和易扩展
写单元测试
语言不限

*/

func PrintAB[T ~int](n T) string {
	var res strings.Builder
	for i := T(1); i <= n; i++ {
		if i%3 == 0 || i%5 == 0 {
			if i%3 == 0 {
				res.WriteString("a")
			}
			if i%5 == 0 {
				res.WriteString("b")
			}
		} else {
			res.WriteString(strconv.FormatInt(int64(i), 10))
		}
	}
	return res.String()
}
