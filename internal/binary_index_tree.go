package internal

// Binary Indexed Tree BIT 树状数组

type BIT []int64

func lowbit(n int) int {
	return n & (-n)
}

func NewBIT(len int) BIT { // BIT 数组序号从1开始。
	return make(BIT, len+1)
}

func (b BIT) QueryN(n int) int64 { // panic if n >= len
	var ans int64
	for ; n > 0; n -= lowbit(n) {
		ans += b[n]
	}
	return ans
}
func (b BIT) Query(l, r int) int64 {
	return b.QueryN(r) - b.QueryN(l)
}

func (b BIT) Update(n int, x int64) {
	for ; n > 0 && n < len(b); n += lowbit(n) {
		b[n] += x
	}
}
