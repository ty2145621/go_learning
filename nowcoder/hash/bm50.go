package hash

/**
BM50 两数之和
知识点
数组
哈希
描述
给出一个整型数组 numbers 和一个目标值 target，请在数组中找出两个加起来等于目标值的数的下标，返回的下标按升序排列。
（注：返回的数组下标从1开始算起，保证target一定可以由数组里面2个数字相加得到）

数据范围：2\leq len(numbers) \leq 10^52≤len(numbers)≤10
5
 ，-10 \leq numbers_i \leq 10^9−10≤numbers
i
​
 ≤10
9
 ，0 \leq target \leq 10^90≤target≤10
9

要求：空间复杂度 O(n)O(n)，时间复杂度 O(nlogn)O(nlogn)
*/

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for k, v := range numbers {
		m[target-v] = k
	}
	for k, v := range numbers {
		if p, ok := m[v]; ok && p != k {
			return []int{k + 1, p + 1}
		}
	}
	return []int{}
}
