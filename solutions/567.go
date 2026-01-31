package main

// проверяем крайний случай, потом вместо hash map можем исполоьзовать просто [26]int. И вместо вычесления количества букв для каждого sliding window
// мы просто вычитаем первое значение и прибовляем последнее + 1 (это эффективней)
func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    charCount := [26]int{}
    charCountCurr := [26]int{}
    for i := range s1 {
        charCount[s1[i] - byte('a')]++
        charCountCurr[s2[i] - byte('a')]++
    }

    l, r := 0, len(s1)
    for r < len(s2) {
        if charCount == charCountCurr {
            return true
        }
        charCountCurr[s2[l] - byte('a')]--
        charCountCurr[s2[r] - byte('a')]++
        l, r = l + 1, r + 1
    }
    if charCount == charCountCurr {
        return true
    } 
    return false
}

