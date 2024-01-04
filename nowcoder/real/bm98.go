package real

/*
*
BM98 螺旋矩阵
题目
题解(154)
讨论(233)
排行
面经new
简单  通过率：30.12%  时间限制：1秒  空间限制：256M
知识点
数组
描述
给定一个m x n大小的矩阵（m行，n列），按螺旋的顺序返回矩阵中的所有元素。

数据范围：0 \le n,m \le 100≤n,m≤10，矩阵中任意元素都满足 |val| \le 100∣val∣≤100
要求：空间复杂度 O(nm)O(nm) ，时间复杂度 O(nm)O(nm)
*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	res := []int{}
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if top > bottom || left > right {
			break
		}
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}

	return res
}

func spiralOrder3(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	res := make([]int, 0)
	res = spiralOrder31(matrix, 0, len(matrix[0])-1, len(matrix)-1, 0, res)

	return res
}

func spiralOrder31(matrix [][]int, up, r, down, l int, res []int) []int {
	if up == down || l == r { // 最后一行、列
		if up == down {
			for i := l; i <= r; i++ {
				res = append(res, matrix[up][i])
			}
		} else {
			for i := up; i <= down; i++ {
				res = append(res, matrix[i][r])
			}
		}
		return res
	}
	if up > down || l > r {
		return res
	}

	for i := l; i < r; i++ {
		res = append(res, matrix[up][i])
	}

	for i := up; i < down; i++ {
		res = append(res, matrix[i][r])
	}

	for i := r; i > l; i-- {
		res = append(res, matrix[down][i])
	}

	for i := down; i > up; i-- {
		res = append(res, matrix[i][l])
	}
	return spiralOrder31(matrix, up+1, r-1, down-1, l+1, res)
}
