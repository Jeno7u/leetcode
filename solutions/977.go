package main

import "fmt"

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
// вместо нахождения центра и двумя поинтерами идти от него постоянно проверяя условия, что поинтеры валидны
// можно изначально задать поинтеры как концы отрезка и идти от них в центр. ТОгда не нужен лишний бинарный поиск и
// не нужно создавать много условий
func sortedSquares(nums []int) []int {
    result := make([]int, len(nums))
	l, r := 0, len(nums) - 1
    for i := len(nums) - 1; i >= 0; i-- {
        if abs(nums[l]) > abs(nums[r]) {
            result[i] = nums[l] * nums[l]
            l++
        } else {
            result[i] = nums[r] * nums[r]
            r--
        }
    }

	return result
}

func main() {
	nums := []int{-1, 2, 2}
	fmt.Println(sortedSquares(nums))
}