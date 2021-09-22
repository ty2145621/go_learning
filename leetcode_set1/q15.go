package code_set1

import (
	"sort"
)

/**
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。

示例：
给定数组 nums = [-1, 0, 1, 2, -1, -4]，
满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	ret := make([][]int, 0)

	if len(nums) < 3 {
		return ret
	}

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k && nums[k] >= 0 {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				if j == i+1 || nums[j] != nums[j-1] {
					ret = append(ret, []int{nums[i], nums[j], nums[k]})
				}
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return ret
}

/* 不排序，很慢
func threeSum(nums []int) [][]int {

	var (
		ret      = make([][]int, 0)
		posNums  = make(map[int]int, 0)
		negaNums = make(map[int]int, 0)
		v0Count  = 0
		m        = make(map[int]int, 0)
		resM     = make(map[string]bool, 0)
	)
	if len(nums) < 3 {
		return ret
	}

	for k, v := range nums {
		if v < 0 {
			negaNums[k] = v
		} else if v > 0 {
			posNums[k] = v
		} else {
			v0Count++
		}
		m[-v] = k
	}

	if v0Count >= 3 {
		ret = append(ret, []int{0, 0, 0})
	}

	for k1, v1 := range posNums {
		for k2, v2 := range negaNums {
			if k3, ok := m[v1+v2]; ok {
				if k3 == k1 || k3 == k2 {
					continue
				}
				res := []int{v1, v2, -v1 - v2}
				sort.Ints(res)
				mapKey := fmt.Sprintf("%v", res)
				if _, ok := resM[mapKey]; ok {
					continue
				}
				ret = append(ret, res)
				resM[mapKey] = true
			}
		}
	}
	return ret
}
*/
