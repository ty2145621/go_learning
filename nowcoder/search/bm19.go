package search

/**
BM19 寻找峰值
给定一个长度为n的数组nums，请你找到峰值并返回其索引。数组可能包含多个峰值，在这种情况下，返回任何一个所在位置即可。
1.峰值元素是指其值严格大于左右相邻值的元素。严格大于即不能有等于
2.假设 nums[-1] = nums[n] = -\infty−∞
3.对于所有有效的 i 都有 nums[i] != nums[i + 1]
4.你可以使用O(logN)的时间复杂度实现此问题吗？

数据范围：
1 \le nums.length \le 2\times 10^5 \1≤nums.length≤2×10
5

-2^{31}<= nums[i] <= 2^{31} - 1−2
31
 <=nums[i]<=2
31
 −1

如输入[2,4,1,2,7,8,4]时，会形成两个山峰，一个是索引为1，峰值为4的山峰，另一个是索引为5，峰值为8的山峰，如下图所示：

示例1
输入：
[2,4,1,2,7,8,4]
复制
返回值：
1
复制
说明：
4和8都是峰值元素，返回4的索引1或者8的索引5都可以
*/

// findPeakElement logn
func findPeakElement(nums []int) int {
	return findChild(nums, 0, len(nums)-1)
}

func findChild(nums []int, low, high int) int {
	if low > high {
		return -1
	}
	if low == high {
		return low
	}

	middle := (low + high) / 2
	// 长度一定>2
	if nums[middle] < nums[middle+1] {
		return findChild(nums, middle+1, high)
	} else if nums[middle] > nums[middle+1] {
		return findChild(nums, low, middle)
	} else {
		return middle
	}
}

func findPeakElement2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}

	if nums[0] > nums[1] {
		return 0
	}

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}

	if nums[len(nums)-2] < nums[len(nums)-1] {
		return len(nums) - 1
	}

	return -1
}
