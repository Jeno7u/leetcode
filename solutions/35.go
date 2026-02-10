package main


// lower bound binary search как раз возвращает индекс для вставки
func searchInsert(nums []int, target int) int {
    l, r := 0, len(nums)
    for l < r {
        mid := (l + r) / 2
        if nums[mid] >= target {
            r = mid
        } else {
            l = mid + 1
        }
    }
    return l
}