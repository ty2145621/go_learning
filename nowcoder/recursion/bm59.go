package recursion

/**
BM59 N皇后问题
题目
题解(71)
讨论(88)
排行
面经new
较难  通过率：50.82%  时间限制：1秒  空间限制：256M
知识点
递归
描述
N 皇后问题是指在 n * n 的棋盘上要摆 n 个皇后，
要求：任何两个皇后不同行，不同列也不在同一条斜线上，
求给一个整数 n ，返回 n 皇后的摆法数。

数据范围: 1 \le n \le 91≤n≤9
要求：空间复杂度 O(1)O(1) ，时间复杂度 O(n!)O(n!)

例如当输入4时，对应的返回值为2，
对应的两种四皇后摆位如下图所示：
*/

func Nqueen(n int) int {
	limit := 1<<n - 1
	return Nqueen1(limit, 0, 0, 0)
}

func Nqueen1(limit, col, right, left int) int {
	if limit == col {
		return 1
	}
	pos := (^(right | left | col)) & limit
	cur, res := 0, 0
	for pos != 0 {
		cur = pos & (-pos)
		res += Nqueen1(limit, col|cur, (right|cur)>>1, (left|cur)<<1)
		pos &= pos - 1
	}
	return res
}

func Nqueen2(n int) int {
	// write code here
	column := make([]bool, n)
	dg := make([]bool, 2*n+1)
	udg := make([]bool, 2*n+1)
	cnt := 0

	Nqueen22(0, n, &cnt, column, dg, udg)
	return cnt
}

func Nqueen22(row, n int, cnt *int, column, dg, udg []bool) {
	if row == n {
		*cnt++
		return
	}

	for i := 0; i < n; i++ {
		if !column[i] && !dg[i-row+n] && !udg[row+i] {
			column[i], dg[i-row+n], udg[row+i] = true, true, true
			Nqueen22(row+1, n, cnt, column, dg, udg)
			column[i], dg[i-row+n], udg[row+i] = false, false, false
		}
	}

}
