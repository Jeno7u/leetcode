package main


func getMostFreq(charFreq [26]int) int {
    result := 0
    for i := range charFreq {
        result = max(result, charFreq[i])
    }
    return result
}

func characterReplacement(s string, k int) int {
    result := 1
    charFreq := [26]int{}
    l, r := 0, 0

    for r < len(s) {
        charFreq[s[r] - 'A']++
        mostFreq := getMostFreq(charFreq)
        if (r - l + 1) - mostFreq > k {
            result = max(result, r - l)
            charFreq[s[l] - 'A']--
            l++
        }
        r++
    }
    return max(result, r - l)
}