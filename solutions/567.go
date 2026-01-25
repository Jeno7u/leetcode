package main

import (
	"slices"
)

// проверяем крайний случай, потом вместо hash map можем исполоьзовать просто [26]int. И вместо вычесления количества букв для каждого sliding window
// мы просто вычитаем первое значение и прибовляем последнее + 1 (это эффективней)
func checkInclusion(s1 string, s2 string) bool {
    if len(s2) < len(s1) {
        return false
    }

    charAmount1 := make([]int, 26)
    charAmount2 := make([]int, 26)
    for i := range s1 {
        charAmount1[int(s1[i]) - int('a')]++
        charAmount2[int(s2[i]) - int('a')]++
    }

    if slices.Equal(charAmount1, charAmount2) {
        return true
    }

    for i := 0; i < len(s2) - len(s1); i++ {
        charAmount2[int(s2[i]) - int('a')]--
        charAmount2[int(s2[i + len(s1)]) - int('a')]++
		if slices.Equal(charAmount1, charAmount2) {
			return true
		}
    }

	return false
}


