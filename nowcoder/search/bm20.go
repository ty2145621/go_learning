package search

import (
	"sort"
)

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


func lowbit(x int) int {
	return x&(-x)
}

func queryN(tree []int, n int) int64 {
	var ans int64
	for n > 0 {
		ans += int64(tree[n])
		n -= lowbit(n)
	}
	return ans
} 

func update(tree []int, n int) {
	for ;n < len(tree); n += lowbit(n) {
		tree[n]++
	}
}

func query(tree []int,a, b int) int64 {
	return queryN(tree, b) - queryN(tree, a)
}
func InversePairs2(data []int) int {
	tree := make([]int, len(data))

	for k, v := range data {
		tree[k] = v
	}

	sort.Ints(tree)
	m := make(map[int]int, 2*len(tree))
	for k, v := range tree {
		m[v] = k+1
	}

	var ans int64
	tree = make([]int, len(data) + 1)
	for i := len(data) - 1; i >= 0; i-- {
		k := m[data[i]]
		ans += queryN(tree, k)
		update(tree, k)
	}

	return int(ans % 1000000007)
}

