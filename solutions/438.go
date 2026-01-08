package main

import (
	"fmt"
	"slices"
)

// исопльзуем hash map или просто [26]int. Делаем sliding window и с каждой итерацией
// убираем появление s[i] из currentFreq и добавляем s[i + len(p)]
func findAnagrams(s string, p string) []int {
    if len(s) < len(p) {
        return []int{}
    }
    
    result := []int{}
	pFreq := make([]int, 26)
	for i := range p {
		pFreq[int(p[i]) - int('a')]++
	}

	currentFreq := make([]int, 26)
	for i := range p{
		currentFreq[int(s[i]) - int('a')]++
	}

	for i := 0; i < len(s) - len(p); i++ {
		if slices.Equal(pFreq, currentFreq) {
			result = append(result, i)
		}
		currentFreq[int(s[i]) - int('a')]--
		currentFreq[int(s[i + len(p)]) - int('a')]++
	}
	if slices.Equal(pFreq, currentFreq) {
		result = append(result, len(s) - len(p))
	}
	return result
}


func main() {
	s := "aaaaaaaaaa"
	p := "aaaaaaaaaaaaa"
	fmt.Println(findAnagrams(s, p))
}