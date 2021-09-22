package leetcode_set1

import (
    "fmt"
    "testing"
)

func TestQ1(t *testing.T) {
    num := []int{2, 7, 11, 15}
    target := 9
    ret := twoSum(num, target)
    fmt.Println(ret)

}

func TestQ3(t *testing.T) {
    sArr := []string{"abcabcbb", "a", "", "abc", "aa", "abccbad", "anviaj"}
    for _, v := range sArr {
        fmt.Println(lengthOfLongestSubstring(v))
    }
}

func TestQ5(t *testing.T) {
    sArr := []string{"ababc", "abcabcbb", "a", "", "abc", "aa", "abccbad", "anviaj"}
    for _, v := range sArr {
        fmt.Println(longestPalindrome(v))
    }
}

func TestQ6(t *testing.T) {
    s := "LDREOEIIECIHNTSGLDREOEIIECIHNTSG"
    fmt.Println(convert(s, 4))
    fmt.Println(convert(s, 5))
    sArr := []string{"ababc", "abcabcbb", "a", "", "abc", "aa", "abccbad", "anviaj"}
    for _, v := range sArr {
        fmt.Println(convert(v, 1))
    }
}

func TestQ7(t *testing.T) {
    iArr := []int{0, 1, -1, 10, -10, 120, -120, 1534236469}
    for _, v := range iArr {
        fmt.Println(reverse(v))
    }
}

func TestQ8(t *testing.T) {
    sArr := []string{"  -42", "a", "+a", "-a", "", "123", "-123", "-120a", "+123aa", "-21534236469aaa", "9223372036854775808"}
    for _, v := range sArr {
        fmt.Println(myAtoi(v))
    }
}

func TestQ9(t *testing.T) {
    iArr := []int{0, 1, -1, 10, 11, 110, 101, 121, 1221, 12321}
    for _, v := range iArr {
        fmt.Println(isPalindrome(v))
    }
}

func TestQ11(t *testing.T) {
    height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
    fmt.Println(maxArea(height))
}

func TestQ12(t *testing.T) {
    iArr := []int{2, 4, 9, 10, 11, 15, 100, 149, 3999}
    for _, v := range iArr {
        fmt.Println(intToRoman(v))
    }
}

func TestQ13(t *testing.T) {
    iArr := []int{2, 4, 9, 10, 11, 15, 100, 149, 3999}
    for _, v := range iArr {
        s := intToRoman(v)
        res := romanToInt(s)
        fmt.Printf("%d:%s:%d\n", v, s, res)
    }
}

func TestQ15(t *testing.T) {
    iArr := [][]int{
        {-1, 0, 1, 2, -1, -4},
        {0, 0, 0},
        {},
        {-1, 0, 1, 0, 0},
        {-1, 0, 1, -1, -1},
        {-2, 0, 1, 1, 1},
        {-2, 0, 0, 0, 2, 2, 2},
        {-2, 0, 0, 2, 2},
    }
    for _, v := range iArr {
        fmt.Printf("%+v:%+v\n", v, threeSum(v))
    }
}

func TestQ16(t *testing.T) {
    iArr := [][]int{
        {-1, 0, 1, 2, -1, -4},
        {-1, 1, 2, -4},
    }
    for _, v := range iArr {
        fmt.Printf("%+v:%+v\n", v, threeSumClosest(v, 1))
    }
}

func TestQ17(t *testing.T) {
    strArr := []string{"234", "2", "222", "23456"}
    for _, v := range strArr {
        fmt.Println(letterCombinations(v))
    }
}

func TestQ18(t *testing.T) {
    intArr := []int{1, 0, -1, 0, -2, 2}
    target := 0
    fmt.Println(fourSum(intArr, target))
}

func TestQ22(t *testing.T) {
    res := generateParenthesis(3)
    for _, v := range res {
        fmt.Println(v)
    }
}

func TestQ28(t *testing.T) {
    hsystack := []string{"ababcaababcaabc", "mississippi", "abc", "abc"}
    needle := []string{"ababcaabc", "pi", "c", "a"}

    for k, _ := range hsystack {
        next := nextPtr(needle[k])
        fmt.Println(next)
        fmt.Println(strStr(hsystack[k], needle[k]))
    }
}

func TestSlice(t *testing.T) {
    s := []int{1, 2, 3, 4, 5}
    fmt.Println(s[:0])
    fmt.Println(s[:1])
    fmt.Println(s[1:])
    fmt.Println(s[5:])

    a1 := s[:1]
    a2 := s[2:]
    s = append(a1, a2...)
    fmt.Println("b", s)
    fmt.Println("s", s)
    d := append(s[:2], s[3:]...)
    fmt.Println(d)
}

func TestQ30(t *testing.T) {
    s := []string{"barfoothefoobarman", "barfoofoobarthefoobarman", "wordgoodgoodgoodbestword", "abababab"}
    needle := [][]string{
        []string{"foo", "bar"},
        []string{"bar", "foo", "the"},
        []string{"word", "good", "best", "good"},
        []string{"ab", "ba"},
    }

    for k, _ := range s {
        fmt.Println(s[k])
        fmt.Println(findSubstring(s[k], needle[k]))
    }
}
