package main

// также важно пробегать по hashSet вместо nums чтобы не смотреть на дубликаты
func longestConsecutive(nums []int) int {
    hashSet := map[int]struct{}{}
    for i := range nums {
        hashSet[nums[i]] = struct{}{}
    }

    result := 0
    for num := range hashSet {
        if _, ok := hashSet[num - 1]; !ok {
            curr := 1
            j := num + 1
            _, ok := hashSet[j]
            for ok {
                curr++
                j++
                _, ok = hashSet[j]
            }
            result = max(result, curr)
        }
    }
    return result
}