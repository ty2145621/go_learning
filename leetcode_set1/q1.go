package leetcode_set1

//给定 nums = [2, 7, 11, 15], target = 9
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
func twoSum(nums []int, target int) []int {
	var m = make(map[int]int, len(nums)) // v => k
	var ret = make([]int, 2)
	for k, v := range nums {
		if ret1, ok := m[v]; ok {
			ret[0] = ret1
			ret[1] = k
			return ret
		}
		m[target-v] = k
	}
	return ret
}
