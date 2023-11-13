package recursion

/**

BM55 没有重复项数字的全排列
题目
题解(79)
讨论(147)
排行
面经new
中等  通过率：59.48%  时间限制：1秒  空间限制：256M
知识点
递归
描述
给出一组数字，返回该组数字的所有排列
例如：
[1,2,3]的所有排列如下
[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2], [3,2,1].
（以数字在数组中的位置靠前为优先级，按字典序排列输出。）

数据范围：数字个数 0 < n \le 60<n≤6
要求：空间复杂度 O(n!)O(n!) ，时间复杂度 O(n!）O(n!）

*/

func permute(num []int) [][]int {
	result := make([][]int, 0)
	result = permute1(0, []int{}, num, result)
	return result
}

func reverse1(begin, pos int, num []int) {
	for i := pos; i > begin; i-- {
		num[i], num[i-1] = num[i-1], num[i]
	}
}
func reverse2(begin, pos int, num []int) {
	for i := begin; i < pos; i++ {
		num[i], num[i+1] = num[i+1], num[i]
	}
}

func permute1(pos int, ans []int, num []int, result [][]int) [][]int {
	if len(ans) == len(num) {
		result = append(result, ans)
	}

	for i := pos; i < len(num); i++ {
		reverse1(pos, i, num)
		result = permute1(pos+1, append(ans, num[pos]), num, result)
		reverse2(pos, i, num)
	}
	return result
}
