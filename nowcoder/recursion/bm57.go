package recursion

/**
BM57 岛屿数量
题目
题解(113)
讨论(168)
排行
面经new
中等  通过率：44.08%  时间限制：1秒  空间限制：256M
知识点
dfs
广度优先搜索(BFS)
搜索
描述
给一个01矩阵，1代表是陆地，0代表海洋， 如果两个1相邻，那么这两个1属于同一个岛。我们只考虑上下左右为相邻。
岛屿: 相邻陆地可以组成一个岛屿（相邻:上下左右） 判断岛屿个数。
例如：
输入
[
[1,1,0,0,0],
[0,1,0,1,1],
[0,0,0,1,1],
[0,0,0,0,0],
[0,0,1,1,1]
]
对应的输出为3
(注：存储的01数据其实是字符'0','1')
*/

func solve(grid [][]byte) int {
	helper := make([][]bool, 0)
	for _, v1 := range grid {
		helper = append(helper, make([]bool, len(v1)))
	}
	count := 0
	solve571(grid, helper, &count)
	return count
}

func solve571(grid [][]byte, helper [][]bool, count *int) bool {
	if len(grid) == 0 {
		return false
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if helper[i][j] {
				continue
			}
			if find571(grid, helper, i, j) {
				*count++
			}
		}
	}
	return true
}

func find571(grid [][]byte, helper [][]bool, i, j int) bool {
	if helper[i][j] {
		return false
	}
	helper[i][j] = true
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return false
	}

	if grid[i][j] != '1' {
		return false
	}

	find571(grid, helper, i+1, j)
	find571(grid, helper, i, j+1)
	find571(grid, helper, i-1, j)
	find571(grid, helper, i, j-1)
	return true
}
