package main

// исопльзуем hash map или просто [26]int. Делаем sliding window и с каждой итерацией
// убираем появление s[i] из currentFreq и добавляем s[i + len(p)]
func findAnagrams(s string, p string) []int {
    if len(s) < len(p) {
        return []int{}
    }
    
    result := []int{}
	pFreq := [26]int{}
	currentFreq := [26]int{}
	for i := range p{
		currentFreq[int(s[i]) - int('a')]++
        pFreq[int(p[i]) - int('a')]++
	}

	for i := 0; i < len(s) - len(p); i++ {
		if pFreq == currentFreq {
			result = append(result, i)
		}
		currentFreq[int(s[i]) - int('a')]--
		currentFreq[int(s[i + len(p)]) - int('a')]++
	}
	if pFreq == currentFreq {
		result = append(result, len(s) - len(p))
	}
	return result
}