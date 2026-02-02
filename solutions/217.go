package main


func containsDuplicate(nums []int) bool {
    hashSet := map[int]struct{}{}

    for i := range nums {
        _, ok := hashSet[nums[i]] 
        if ok {
            return true
        }
        hashSet[nums[i]] = struct{}{}
    }
    return false
}