package main

// ставим l и r на концах для максимизации объема.
// Потом просто сдвигаем поинтер, где самая маленькая высота
func maxArea(height []int) int {
    result := 0
    l, r := 0, len(height) - 1
    for l < r {
        currVolume := min(height[l], height[r]) * (r - l)
        result = max(result, currVolume)

        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }
    return result
}