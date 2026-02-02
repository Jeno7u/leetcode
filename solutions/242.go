package main

// вместо hashMap быстрее использовать [26]int
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    charFreq := [26]int{}
    for i := range s {
        charFreq[s[i] - 'a']++
    }

    for i := range t {
        val := charFreq[t[i] - 'a']
        if val <= 0 {
            return false
        }
        charFreq[t[i] - 'a']--
    }
    return true
}