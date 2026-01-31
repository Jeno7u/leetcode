package main


func isValid(s string) bool {
    mapOpen := map[byte]struct{}{'(': struct{}{}, '[': struct{}{}, '{': struct{}{}}
    mapClosed := map[byte]byte{')': '(', ']': '[', '}': '{'}
    stack := []byte{}

    for i := 0; i < len(s); i++ {
        _, ok := mapOpen[s[i]]
        if ok {
            stack = append(stack, s[i])
        } else {
            val, _ := mapClosed[s[i]]
            if len(stack) == 0 || val != stack[len(stack) - 1] {
                return false
            }
            stack = stack[:len(stack) - 1]
        }
    }

    if len(stack) != 0 {
        return false
    }
    return true
}