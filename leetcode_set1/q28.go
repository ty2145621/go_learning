package leetcode_set1

/**
28. 实现 strStr()
实现 strStr() 函数。
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
示例 1:
输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:
输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:
当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
*/

//next[j]沒有设置为-1, 跟带有-1的区别在于next ptr是否指向的是下一个匹配pos还是需要n++才能指向下一个匹配pos
// KMP算法： next数组即为 找到最长的前缀后缀匹配长度
func nextPtr(needle string) (next []int) {
    next = make([]int, len(needle))

    // aabaacaab
    k := 0
    for i := 2; i < len(needle); i++ {
        for k > 0 && needle[i-1] != needle[k] {
            k = next[k]
        }
        if needle[i-1] == needle[k] {
            k++
        }
        next[i] = k
    }

    return next
}

func strStr(haystack string, needle string) int {
    if len(needle) == 0 {
        return 0
    }
    next := nextPtr(needle)

    n := 0
    for h := 0; h < len(haystack); h++ {
        for n > 0 && haystack[h] != needle[n] {
            n = next[n]
        }
        if haystack[h] == needle[n] {
            n++
        }
        if n == len(needle) {
            return h - n + 1
        }
    }

    return -1
}

/*// KMP算法： next数组即为 找到最长的前缀后缀匹配长度
func nextPtr(needle string) (next []int) {
    next = make([]int, len(needle))
    next[0] = -1
    // aabaacaab
    k := -1
    for i := 1; i < len(needle); i++ {
        for k >= 0 && needle[i] != needle[k+1] {
            k = next[k]
        }
        if needle[i] == needle[k+1] {
            k++
        }
        next[i] = k
    }

    return next
}

func strStr(haystack string, needle string) int {
    if len(needle) == 0 {
        return 0
    }
    next := nextPtr(needle)

    n := -1
    for h := 0; h < len(haystack); h++ {
        for n >= 0 && haystack[h] != needle[n+1] {
            n = next[n]
        }
        if haystack[h] == needle[n+1] {
            n++
        }
        if n == len(needle)-1 {
            return h - n
        }
    }

    return -1
}
*/
