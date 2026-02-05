package main

import (
	"cmp"
	"slices"
)

// если нет пересечений, то просто забиваем и идем дальше
// если есть, то мы убираем тот который заканчивается позже (можно даже так без удаления физического)
func eraseOverlapIntervals(intervals [][]int) int {
    slices.SortFunc(intervals, func(a, b []int) int {return cmp.Compare(a[0], b[0])})
    result := 0

    prevEnd := intervals[0][1]
    for _, val := range intervals[1:] {
        start, end := val[0], val[1]
        if start >= prevEnd {
            prevEnd = end
        } else {
            result++
            prevEnd = min(prevEnd, end)
        }
    }
    return result
}