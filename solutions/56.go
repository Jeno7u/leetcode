package main

import (
	"cmp"
	"slices"
)


func anotherMerge(intervals [][]int) [][]int {
    slices.SortFunc(intervals, func(a, b []int) int { return cmp.Compare(a[0], b[0]) })

	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		// добавить если новый отрезок
		if intervals[i][0] > result[len(result) - 1][1]{
			result = append(result, intervals[i])
		// обьеденить с помощью нахождения конца отрезка 
		} else {
			result[len(result) - 1][1] = max(result[len(result) - 1][1], intervals[i][1])
		}
	}
	return result
}
