package main

import (
	"cmp"
	"slices"
)

// используется сортировка для вставки элемента
// можно вообще сделать и кастомную сортировку для вставки элемента
// или добавлять в result все элементы пока не дойдем до пересечения с newInterval
// тогда мы объединяем и вставляем оставшиеся элементы
func insert(intervals [][]int, newInterval []int) [][]int {
    intervals = append(intervals, newInterval)
    slices.SortFunc(intervals, func(a, b []int) int {
        return cmp.Compare(a[0], b[0])
    })

    result := [][]int{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        lastIdx := len(result) - 1
        if result[lastIdx][1] >= intervals[i][0] {
            result[lastIdx][1] = max(result[lastIdx][1], intervals[i][1])
        } else {
            result = append(result, intervals[i])
        }
    }
    return result
}