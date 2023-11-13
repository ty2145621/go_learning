package real

/*
*
BM99 顺时针旋转矩阵
题目
题解(102)
讨论(275)
排行
面经new
中等  通过率：53.34%  时间限制：1秒  空间限制：256M
知识点
数组
基础数学
描述
有一个NxN整数矩阵，请编写一个算法，将矩阵顺时针旋转90度。

给定一个NxN的矩阵，和矩阵的阶数N,请返回旋转后的NxN矩阵。
*/
func rotateMatrix(mat [][]int, n int) [][]int {
	n = n - 1
	for i := 0; i < (n+1)/2; i++ {
		for j := 0; j < (n+2)/2; j++ {
			mat[i][j], mat[j][n-i], mat[n-i][n-j], mat[n-j][i] = mat[n-j][i], mat[i][j], mat[j][n-i], mat[n-i][n-j]
		}
	}
	return mat
}
