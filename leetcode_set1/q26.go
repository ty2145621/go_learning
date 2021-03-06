package leetcode_set1

/**
26. 删除排序数组中的重复项
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

示例 1:
给定数组 nums = [1,1,2],
函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
你不需要考虑数组中超出新长度后面的元素。
*/
func removeDuplicates(nums []int) int {
    k := 0
    cur := 0
    for k < len(nums) {
        for k+1 < len(nums) && nums[k] == nums[k+1] {
            k++
        }
        if cur != k {
            nums[cur] = nums[k]
        }
        cur++
        k++
    }
    return cur
}
/*
func removeDuplicates(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }
    cur := 0
    for k := 1; k < len(nums); k++ {
        if nums[k] != nums[cur] {
            cur++
            nums[cur] = nums[k]
        }
    }
    return cur+1
}
*/