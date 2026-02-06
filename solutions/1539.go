package main


func findKthPositive(arr []int, k int) int {
    set := map[int]struct{}{}
    for i := range arr {
        set[arr[i]] = struct{}{}
    }

    length := len(arr) + k + 1
    for i := 1; i < length; i++ {
        if _, ok := set[i]; !ok {
            k--
        }
        if k == 0 {
            return i
        }
    }
    return -1
}