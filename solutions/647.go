package main


func countSubstrings(s string) int {
    result := 0

    for i := range s {
        l, r := i, i // это центр нашего палиндрома нечетного
        for l >= 0 && r < len(s) && s[l] == s[r] {
            result++
            l--
            r++
        }

        l, r = i, i + 1 // это центр нашего палиндрома четного
        for l >= 0 && r < len(s) && s[l] == s[r] {
            result++
            l--
            r++
        }
    }
    return result
}