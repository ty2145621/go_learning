package search

/**
BM20 数组中的逆序对
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组,求出这个数组中的逆序对的总数P。并将P对1000000007取模的结果输出。 即输出P mod 1000000007

数据范围：  对于 50\%50% 的数据, size\leq 10^4size≤10
4

对于 100\%100% 的数据, size\leq 10^5size≤10
5

数组中所有数字的值满足 0 \le val \le 10^90≤val≤10
9


要求：空间复杂度 O(n)O(n)，时间复杂度 O(nlogn)O(nlogn)
输入描述：
题目保证输入的数组中没有的相同的数字
示例1
输入：
[1,2,3,4,5,6,7,0]
复制
返回值：
7
复制
示例2
输入：
[1,2,3]
复制
返回值：
0
*/

var res []int

func InversePairs(data []int) int {
	res = make([]int, len(data))
	mid := len(data) / 2

	count := mergeSort(data, 0, mid, len(data))
	return count % 1000000007
}

func mergeSort(nums []int, start, mid, end int) (count int) {
	var c1, c2 int
	if mid-start > 1 {
		m := (mid + start) / 2
		c1 = mergeSort(nums, start, m, mid)
	}

	if end-mid > 1 {
		m := (end + mid) / 2
		c2 = mergeSort(nums, mid, m, end)
	}

	i, j, p := start, mid, 0
	count = c1 + c2
	for i < mid && j < end {
		if nums[i] > nums[j] {
			res[p] = nums[j]
			j++
			p++
			count += mid - i
		} else if nums[i] < nums[j] {
			res[p] = nums[i]
			i++
			p++
		} else {
			// fmt.Println("no equal num")
		}
	}

	for i < mid {
		res[p] = nums[i]
		i++
		p++
	}
	for j < end {
		res[p] = nums[j]
		j++
		p++
	}

	for i := 0; i < end-start; i++ {
		nums[start+i] = res[i]
	}

	return count
}
