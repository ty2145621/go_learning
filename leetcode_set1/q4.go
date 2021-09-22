package leetcode_set1

//4. 寻找两个正序数组的中位数
//这个题目可以归结到寻找第k小(大)元素问题，思路可以总结如下：取两个数组中的第k/2个元素进行比较，如果数组1的元素小于数组2的元素，则说明数组1中的前k/2个元素不可能成为第k个元素的候选，所以将数组1中的前k/2个元素去掉，组成新数组和数组2求第k-k/2小的元素，因为我们把前k/2个元素去掉了，所以相应的k值也应该减小。另外就是注意处理一些边界条件问题，比如某一个数组可能为空或者k为1的情况。
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	if len1 == 0 && len2 == 0 {
		return 0
	}
	left, right := (len1+len2+1)/2, (len1+len2+2)/2
	return (findKthElement(nums1, nums2, left) + findKthElement(nums1, nums2, right)) / 2
}

func findKthElement(nums1 []int, nums2 []int, k int) float64 {
	len1, len2 := len(nums1), len(nums2)
	if len1 > len2 {
		return findKthElement(nums2, nums1, k)
	}

	if len1 == 0 {
		return float64(nums2[k-1])
	}

	if k == 1 {
		return float64(getMin(nums1[0], nums2[0]))
	}

	p1 := getMin(k/2, len1)-1
	p2 := getMin(k/2, len2)-1
	if nums1[p1] < nums2[p2] {
		return findKthElement(nums1[p1+1:len1], nums2, k-p1-1)
	} else {
		return findKthElement(nums1, nums2[p2+1:len2], k-p2-1)
	}
}

func getMin(k1, k2 int) int {
	if k1 > k2 {
		return k2
	}
	return k1
}
