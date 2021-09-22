package leetcode_set1

/**
22. 括号生成
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
示例：
输入：n = 3
输出：[
       "((()))",
       "(()())",
       "(())()",
       "()(())",
       "()()()"
     ]
*/
func generateParenthesis(n int) []string {
    //return generateParenthesisRecursion(n)
    return generateWithDongtaiguihua(n)
}

func generateWithDongtaiguihua(n int) []string {
    var (
        res = make([][]string, n+1)
    )

    if n == 0 {
        return []string{}
    }

    res[0] = append(res[0], "")
    res[1] = append(res[1], "()")

    for i := 2; i <= n; i++ {
        for j := 0; j < i; j++ {
            for _, m := range res[j] {
                for _, n := range res[i-j-1] {
                    res[i] = append(res[i], "("+m+")"+n)
                }
            }
        }
    }
    return res[n]
}

func generateParenthesisRecursion(n int) []string {
    var resArr = make([]string, 0)
    geneRecur(n, 0, 0, "", &resArr)

    return resArr
}

// n
func geneRecur(n int, left int, right int, res string, resArr *[]string) {
    if left > n || left < right {
        return
    }
    if left == right && right == n {
        *resArr = append(*resArr, res)
    }

    geneRecur(n, left+1, right, res+"(", resArr)
    geneRecur(n, left, right+1, res+")", resArr)
}
