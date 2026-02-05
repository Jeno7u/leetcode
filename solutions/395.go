package main

// так как мы не знааем когда нам сужать и расширять точно окно (скорее только сужать)
// мы проходимся по currentUniqCount и добираем наш sliding window до такого размера.
// Если перескакиваем, то сужаем
func longestSubstring(s string, k int) int {
    maxUniqueCharCount := getUniqueCharCount(s)
    freqMap := [26]int{}
    result := 0

    for currentUniqueCount := 1; currentUniqueCount <= maxUniqueCharCount; currentUniqueCount++ {
        clearArray(&freqMap)
        uniqueCount, countAtLeastK := 0, 0
        windowStart, windowEnd := 0, 0

        for windowEnd < len(s) {
            if uniqueCount <= currentUniqueCount {
                idx := s[windowEnd] - 'a'
                if freqMap[idx] == 0 {
                    uniqueCount++
                }
                freqMap[idx]++

                if freqMap[idx] == k {
                    countAtLeastK++
                }
                windowEnd++
            } else {
                idx := s[windowStart] - 'a'
                if freqMap[idx] == k {
                    countAtLeastK--
                }
                freqMap[idx]--

                if freqMap[idx] == 0 {
                    uniqueCount--
                }
                windowStart++
            }

            if uniqueCount == currentUniqueCount && uniqueCount == countAtLeastK {
                result = max(result, windowEnd - windowStart)
            }
        }
    }
    return result
}


func getUniqueCharCount(s string) int {
    chars := [26]bool{}
    uniqueCharCount := 0
    for i := range s {
        if !chars[s[i] - 'a'] {
            uniqueCharCount++
        }
        chars[s[i] - 'a'] = true
    }
    return uniqueCharCount
}

func clearArray(arr *[26]int) {
    for i := range *arr {
        (*arr)[i] = 0
    }
}