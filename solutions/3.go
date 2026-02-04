package main


func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }
    result := 1
    currChars := map[byte]int{}
    l, r := 0, 0
    for r < len(s) {
        val, ok := currChars[s[r]]
        if !ok {
            currChars[s[r]] = r
        } else {
            if l <= val && val <= r {
                l = val + 1
            }
            currChars[s[r]] = r
        }
        result = max(result, r - l + 1)
        r++
    }
    return max(result, r - l)
}