package main


func reverseWords(s string) string {
    result := []byte{}
    l, r := 0, 0
    for r < len(s) {
        if s[r] == ' ' {
            for i := r - 1; i >= l; i-- {
                result = append(result, s[i])
            }
            result = append(result, s[r])
            l = r + 1
        }
        r++
    }
    for i := r - 1; i >= l; i-- {
        result = append(result, s[i])
    }
    return string(result)
}