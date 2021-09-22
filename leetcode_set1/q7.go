package leetcode_set1

// 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//
//示例 1:
//输入: 120
//输出: 21
//
// 输入: -123
// 输出: -321

const INT_MAX = int32(^uint32(0) >> 1)
const INT_MIN = ^INT_MAX

func reverse(x int) int {
	var (
		ret     = 0
		negative bool
	)
	if x < 0 {
		negative = true
		x *= -1
	}

	if x == 0 {
		return x
	}

	for x > 0 {
		lastNum := x % 10
		ret = ret*10 + lastNum
		x /= 10
	}
	if negative {
		ret *= -1
	}

	if ret > int(INT_MAX) || ret < int(INT_MIN) {
		ret = 0
	}

	return ret
}
