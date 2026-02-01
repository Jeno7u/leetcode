package main

// Храним в LIS по индексу самую длинную возможную послежовательность если добавить nums[i].
// Обновляем LIS только если новое значение меньше минимума в последовательности (< nums[j])
func lengthOfLIS(nums []int) int {
    LIS := make([]int, len(nums))
    for i := range LIS {
        LIS[i] = 1
    }

    for i := len(nums) - 1; i >= 0; i-- {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] < nums[j] {
                LIS[i] = max(LIS[i], 1 + LIS[j])
            }
        }
    }
    
    maximum := 0
    for i := range LIS {
        maximum = max(maximum, LIS[i])
    }
    return maximum
}