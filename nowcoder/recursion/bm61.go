package recursion

import (
	"fmt"
)

/**
BM61 矩阵最长递增路径
题目
题解(55)
讨论(57)
排行
面经new
中等  通过率：43.19%  时间限制：3秒  空间限制：256M
知识点
dfs
动态规划
图
描述
给定一个 n 行 m 列矩阵 matrix ，矩阵内所有数均为非负整数。 你需要在矩阵中找到一条最长路径，使这条路径上的元素是递增的。并输出这条最长路径的长度。
这个路径必须满足以下条件：

1. 对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外。
2. 你不能走重复的单元格。即每个格子最多只能走一次。

数据范围：1 \le n,m \le 10001≤n,m≤1000，0 \le matrix[i][j] \le 10000≤matrix[i][j]≤1000
进阶：空间复杂度 O(nm)O(nm) ，时间复杂度 O(nm)O(nm)

例如：当输入为[[1,2,3],[4,5,6],[7,8,9]]时，对应的输出为5，
其中的一条最长递增路径如下图所示：
*/

func solve61(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	helper := make([][]int, 0)
	for _, v := range matrix {
		helper = append(helper, make([]int, len(v)))
	}

	for i, v1 := range helper {
		for j, _ := range v1 {
			solve611(matrix, helper, i, j)
		}
	}
	fmt.Println(helper)
	m := 0
	for _, v1 := range helper {
		for _, v2 := range v1 {
			if m < v2 {
				m = v2
			}
		}
	}
	return m
}

// helper 记录每个开头的 最大长度
func solve611(matrix [][]int, helper [][]int, i, j int) int {
	if helper[i][j] != 0 {
		return helper[i][j]
	}
	helper[i][j] = 1

	m1, m2, m3, m4 := 0, 0, 0, 0
	if i < len(matrix)-2 && matrix[i+1][j] > matrix[i][j] {
		m1 = solve611(matrix, helper, i+1, j)
	}
	if j < len(matrix[0])-2 && matrix[i][j+1] > matrix[i][j] {
		m2 = solve611(matrix, helper, i, j+1)
	}
	if i > 0 && matrix[i-1][j] > matrix[i][j] {
		m3 = solve611(matrix, helper, i-1, j)
	}
	if j > 0 && matrix[i][j-1] > matrix[i][j] {
		m4 = solve611(matrix, helper, i, j-1)
	}
	m := getMax61(m1, m2, m3, m4) + 1
	helper[i][j] = m
	return m
}

func getMax61(a ...int) int {
	m := a[0]
	for _, v := range a {
		if m < v {
			m = v
		}
	}
	return m
}
