package leetcode_set1

import "sort"

/**
4sum
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
答案中不可以包含重复的四元组。

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/

func getNext(a []int, begin int) int {
    begin = begin + 1
    for begin < len(a) && a[begin] == a[begin-1] {
        begin++
    }
    return begin
}

func getBefore(a []int, begin int, end int) int {
    end = end - 1
    for end > begin && a[end] == a[end+1] {
        end--
    }
    return end
}

func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    var (
        ret = make([][]int, 0)
    )
    if len(nums) < 4 {
        return ret
    }

    for a1 := 0; a1 < len(nums); a1 = getNext(nums, a1) {
        for i := a1 + 1; i < len(nums)-2; i = getNext(nums, i) {
            j := i + 1
            k := len(nums) - 1
            for j < k {
                sum := nums[a1] + nums[i] + nums[j] + nums[k]
                if sum == target {
                    ret = append(ret, []int{nums[a1], nums[i], nums[j], nums[k]})
                    j = getNext(nums, j)
                    k = getBefore(nums, j, k)
                } else if sum < target {
                    j = getNext(nums, j)
                } else {
                    k = getBefore(nums, j, k)
                }
            }
        }
    }
    return ret
}
