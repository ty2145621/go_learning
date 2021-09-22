package code_set1

/**
20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:
*/

/*func isValid(s string) bool {
    var (
        arr = make([]uint8, 0)
    )
    if len(s)%2 != 0 {
        return false
    }
    for _, v := range []byte(s) {
        if v == '(' || v == '[' || v == '{' {
            arr = append(arr, v)
        } else if len(arr) > 0 && (
            (v == ')' && arr[len(arr)-1] == '(') ||
                (v == ']' && arr[len(arr)-1] == '[') ||
                (v == '}' && arr[len(arr)-1] == '{')) {
            arr = arr[:len(arr)-1]
        } else {
            return false
        }
    }

    if len(arr) > 0 {
        return false
    }
    return true
}*/

func isValid(s string) bool {
    stack := make([]byte,0)
    charMap := map[byte]byte{
        '(':')',
        '{':'}',
        '[':']',
    }
    for i:=0; i < len(s); i++ {
        c := s[i]
        length := len(stack)
        if length == 0 {
            stack = append(stack, c)
            continue
        }
        if charMap[stack[length-1]] == c {
            stack = stack[:length-1]
        }else{
            stack = append(stack, c)
        }
    }
    return len(stack) == 0
}
