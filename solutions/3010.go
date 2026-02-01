package main

// первое значение мы берем всегда. А остальные два это минимальные значения в списке.
// Так как мы можем по сути брать любые значения в списке (я про 2 и 3), то мы берем просто
// самые маленькие (по условию)
func minimumCost(nums []int) int {
    result := nums[0]
    min1, min2 := 51, 51

    for i := 1; i < len(nums); i++ {
        if nums[i] < min1 {
            min1, min2 = nums[i], min1
        } else if nums[i] < min2 {
            min2 = nums[i]
        }
    }
    return result + min1 + min2
}