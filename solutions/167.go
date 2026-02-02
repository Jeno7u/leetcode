package main


func binarySearch(numbers []int, l, r, target int) int {
    for l < r {
        mid := (l + r) / 2
        if numbers[mid] == target {
            return mid
        }
        if numbers[mid] < target {
            l = mid + 1
        } else {
            r = mid
        }
    }
    return -1
}

// или берем два поинтера на концах и считаем сумму. Если больше чем таргет,
// то r--, иначе l++
func twoSum(numbers []int, target int) []int {
    for i := range numbers {
        missing := target - numbers[i]
        idx := binarySearch(numbers, i, len(numbers), missing)
        if idx != -1 {
            return []int{i + 1, idx + 1}
        }
    }
    return []int{}
}