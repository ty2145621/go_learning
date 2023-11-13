package recursion

/**
BM60 括号生成
题目
题解(81)
讨论(137)
排行
面经new
中等  通过率：57.05%  时间限制：1秒  空间限制：256M
知识点
递归
描述
给出n对括号，请编写一个函数来生成所有的由n对括号组成的合法组合。
例如，给出n=3，解集为：
"((()))", "(()())", "(())()", "()()()", "()(())"

数据范围：0 \le n \le 100≤n≤10
要求：空间复杂度 O(n)O(n)，时间复杂度 O(2^n)O(2
n
 )
*/

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	return generateParenthesis1(n, n, []byte{}, res)
}

// generateParenthesis1
// nl 左括号剩余
// nr 右括号剩余
func generateParenthesis1(nl, nr int, bArr []byte, res []string) []string {
	if nl == 0 && nr == 0 {
		res = append(res, string(bArr))
	}
	if nl < 0 || nr < 0 || nl > nr {
		return res
	}
	if nl < nr {
		res = generateParenthesis1(nl, nr-1, append(bArr, ')'), res)
	}
	res = generateParenthesis1(nl-1, nr, append(bArr, '('), res)

	return res
}
