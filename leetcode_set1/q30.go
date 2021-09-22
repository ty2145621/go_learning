package code_set1

import (
    "fmt"
    "strings"
)

/**
30. 串联所有单词的子串
给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

示例 1：

输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：

输入：
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
输出：[]
*/

// 这道题的解法，都有问题
// aa aa
// aaaaaaaaaa

func HasPrefixInArr(s string, l int, h map[string]int) bool {
    n := make(map[string]int)
    for k, v := range h {
        n[k] = v
    }

    for i := 0; i+l <= len(s); i += l {
        if n[s[i:i+l]] <= 0 {
            return false
        }
        n[s[i:i+l]] --
    }
    return true
}

func checkMapZero(ha map[uint8]int) bool {
    for _, v := range ha {
        if v != 0 {
            return false
        }
    }
    return true
}

func findSubstring(s string, words []string) []int {
    res := make([]int, 0)
    if len(words) == 0 {
        return res
    }

    l := len(words[0])
    totalL := l * len(words)
    if len(s) < totalL {
        return res
    }

    h := make(map[string]int, 0)
    for _, v := range words {
        h[v] = h[v] + 1
    }
    fmt.Println(h)

    ha := make(map[uint8]int, 0)
    for _, v := range []byte(strings.Join(words, "")) {
        ha[v] = ha[v] + 1
    }

    for p := 0; p < totalL; p++ {
        ha[s[p]] = ha[s[p]] - 1
    }

    if checkMapZero(ha) {
        if HasPrefixInArr(s[:totalL], l, h) {
            res = append(res, 0)
        }
    }

    for p := 1; p+totalL-1 < len(s); p++ {
        ha[s[p-1]] ++
        ha[s[p+totalL-1]] --
        if ha[s[p-1]] != 0 || ha[s[p+totalL-1]] != 0 {
            continue
        }
        check := checkMapZero(ha)
        if !check {
            continue
        }

        exist := HasPrefixInArr(s[p:p+totalL], l, h)
        if exist {
            res = append(res, p)
            continue
        }

    }

    return res
}
