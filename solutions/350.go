package main

// берем один список и считаем кол-во определенных чисел. Потом проходимся по второму и смотрим есть ли такое число в первом.
// Если есть уменьшаем счетчик кол-ва такого числа, добавляем в рещультат и идем дальше
func intersect(nums1 []int, nums2 []int) []int {  
    numCount2 := map[int]int{}
    for i := range nums2 {
        numCount2[nums2[i]]++
    }

    result := []int{}
    for i := range nums1 {
        if val, _ := numCount2[nums1[i]]; val > 0 {
            result = append(result, nums1[i])
            numCount2[nums1[i]]--
        }
    }
    return result
}
