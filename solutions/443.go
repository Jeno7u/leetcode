package main

import "strconv"


func compress(chars []byte) int {
    idx, l, r := 0, 0, 0

    for r < len(chars) {
        if r == len(chars) - 1 || chars[r] != chars[r + 1] {
            chars[idx] = chars[r]
            if r - l + 1 == 1 {
                idx++
            } else {
                idx++
                convertedCount := strconv.Itoa(r - l + 1)
                for i := 0; i < len(convertedCount); i++ {
                    chars[idx] = convertedCount[i]
                    idx++
                }

            }
            l = r + 1
        }
        r++
    }
    return idx
}