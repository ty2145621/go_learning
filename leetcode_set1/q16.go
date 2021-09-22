package code_set1

import (
	"sort"
)

/**
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

示例：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。

提示：
3 <= nums.length <= 10^3
-10^3 <= nums[i] <= 10^3
-10^4 <= target <= 10^4

*/
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var closet = 10000000

	if len(nums) < 3 {
		return closet
	}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if abs(sum-target) < abs(closet-target) {
				closet = sum
			}
			if sum == target {
				return target
			} else if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	return closet
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
