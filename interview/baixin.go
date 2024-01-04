package interview

// quickSort 快速排序
func quickSort(num []int) {
	if len(num) <= 1 {
		return
	}
	l, r := 0, len(num)-1
	p := num[0]

	for l < r {
		for l < r && num[r] >= p {
			r--
		}
		num[l] = num[r]
		for l < r && num[l] <= p {
			l++
		}
		num[r] = num[l]
	}
	num[l] = p

	quickSort(num[:l])
	quickSort(num[l+1:])
}

// findAll 找出数组的所有组合
func findAll(num []int) [][]int {
	// 1,2,3,4,5
	p := make([][]int, 0)
	p = append(p, []int{1})
	for i := 1; i < len(num); i++ {
		for _, v := range p {
			p = append(p, append(v, num[i]))
		}
	}
	return p
}
