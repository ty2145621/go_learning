package recursion

import (
	"sort"
)

/**

BM56 有重复项数字的全排列
面经new
中等  通过率：42.82%  时间限制：1秒  空间限制：256M
知识点
递归
描述
给出一组可能包含重复项的数字，返回该组数字的所有排列。结果以字典序升序排列。

数据范围： 0 < n \le 80<n≤8 ，数组中的值满足 -1 \le val \le 5−1≤val≤5
要求：空间复杂度 O(n!)O(n!)，时间复杂度 O(n!)O(n!)
*/

// TODO:error
func permuteUnique(num []int) [][]int {
	sort.Ints(num)
	result := make([][]int, 0)
	result = permuteUnique1(num, result)
	return result
}

func permuteUnique1(num []int, result [][]int) [][]int {
	ok := true
	for ok {
		dst := make([]int, len(num))
		copy(dst, num)
		result = append(result, dst)
		ok, _ = nextPermutation(num)
	}
	return result
}

func nextPermutation(num []int) (bool, []int) {
	i := len(num) - 2
	for ; i >= 0; i-- {
		if num[i] < num[i+1] {
			break
		}
	}
	if i < 0 {
		return false, []int{}
	}
	j := len(num) - 1
	for ; j >= 0; j-- {
		if num[i] < num[j] {
			break
		}
	}

	num[i], num[j] = num[j], num[i]
	i++
	for j := len(num) - 1; i < j; {
		num[i], num[j] = num[j], num[i]
		i++
		j--
	}
	return true, num
}
