package main

// движемся вправо через r до момоента пока zerosCount <= k. Если больше, то мы двигаем левый поинтер пока nums[l] не будет равен 0, постоянно сдвигая l
// ну и result у нас это просто расстояние между l и r
func longestOnes(nums []int, k int) int {
    l, zerosCount, result := 0, 0, 0
    for r := 0; r < len(nums); r++ {
        if nums[r] == 0 {
            zerosCount++
        }
        for zerosCount > k {
            if nums[l] == 0 {
                zerosCount--
            }
            l++
        }
        result = max(result, r - l + 1)
    }
    return result
}

