package main

import (
	"fmt"
	"maps"
	"slices"
)

// несмотря на то что казалось бы нужно хешировать словарь с частотами букв, можно использовать просто [26]int{} так как используются только
// маленькие буквы английские, а в оставльном не так сложно
func groupAnagrams(strs []string) [][]string {
    freq := map[[26]int][len(strs)]string{}
	for _, word := range strs {
		wordFreq := [26]int{}
		for _, char := range word {
			wordFreq[int(char) - int('a')]++
		}
		freq[wordFreq] = append(freq[wordFreq], word)
	}
	fmt.Println(freq)

	return slices.Collect(maps.Values(freq))
}

func main() {
	strs := []string{"eat","tea","tan","ate","nat","bat"}
	fmt.Println(groupAnagrams(strs))
}