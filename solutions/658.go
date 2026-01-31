package main


func findClosestElements(arr []int, k int, x int) []int {
    l, r := 0, len(arr) - k
    for l < r {
        mid := (l + r) / 2
        if x - arr[mid] <= arr[mid+k] - x {
            r = mid
        } else {
            l = mid + 1
        }
    }
    return arr[l:r+k]
}