package main

// вместо того чтобы менять местами два элемента, проще просто заменять по порядку
func removeElement(nums []int, val int) int {
    idx := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[idx] = nums[i]
            idx++
        }
    }
    return idx
}