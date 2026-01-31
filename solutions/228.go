package main

import (
	"fmt"
	"strconv"
)


func summaryRanges(nums []int) []string {
    result := []string{}

    l, r := 0, 0
    for r < len(nums) - 1 {
        if nums[r+1] - nums[r] == 1 {
            r++
        } else {
            if l != r {
                result = append(result, fmt.Sprintf("%d->%d", nums[l], nums[r]))
            } else {
                result = append(result, strconv.Itoa(nums[r]))
            }
            r = r + 1
            l = r
        }
    }
    if l != r {
        result = append(result, fmt.Sprintf("%d->%d", nums[l], nums[r]))
    } else {
        result = append(result, strconv.Itoa(nums[r]))
    }

    return result
}