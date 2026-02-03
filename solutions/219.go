package main


func containsNearbyDuplicate(nums []int, k int) bool {
    numIdx := map[int]int{}
    for i := range nums {
        index, ok := numIdx[nums[i]]
        if ok && i - index <= k {
            return true
        }
        numIdx[nums[i]] = i
    }
    return false
}