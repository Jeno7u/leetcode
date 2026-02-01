package main


func intersection(nums1 []int, nums2 []int) []int {
    hashSet := map[int]struct{}{}

    for i := range nums1 {
        hashSet[nums1[i]] = struct{}{}
    }

    result := []int{}
    for i := range nums2 {
        if _, ok := hashSet[nums2[i]]; ok {
            result = append(result, nums2[i])
            delete(hashSet, nums2[i])
        }
    }
    return result
}