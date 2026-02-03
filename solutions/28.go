package main


func strStr(haystack string, needle string) int {
    if len(needle) > len(haystack) {
        return -1
    }

    i := 0
    j := 0
    for i < len(haystack) - len(needle) + 1 {
        for haystack[i] == needle[j] {
            i++
            j++
            if j == len(needle) {
                return i - len(needle)
            }
        }
        i -= j
		j = 0
        i++
    }
    return -1
}   
