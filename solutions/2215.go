package main


func findDifference(nums1 []int, nums2 []int) [][]int {
    result := [][]int{{}, {}}

    hashSet2 := map[int]struct{}{}
    for i := range nums2 {
        hashSet2[nums2[i]] = struct{}{}
    }

    hashSet1 := map[int]struct{}{}
    for i := range nums1 {
        hashSet1[nums1[i]] = struct{}{}
    }

    for val, _ := range hashSet1 {
        _, ok := hashSet2[val]
        if !ok {
            result[0] = append(result[0], val)
        }
    }

    for val, _ := range hashSet2 {
        _, ok := hashSet1[val]
        if !ok {
            result[1] = append(result[1], val)
        }
    }

    return result
}