package main

import "fmt"

// в версии .py мы проходимся по string t в поисках совпадений по s. А здесь мы проходимся по string s в поисках совпадений в t
func isSubsequence(s string, t string) bool {
    if len(t) < len(s) {
        return false
    }

    idx_t := 0
    for i := 0; i < len(s); i++ {
		fmt.Println(i, idx_t)
		found := false
        // go through string t to find occurance of s[i]
        for idx_t < len(t) {
            if s[i] == t[idx_t] {
				found = true
				idx_t++
                break
            }
            idx_t++
        }

        // if it is not last in s but idx_t is already done or if not founded -> return false
        if (i != len(s) - 1 && idx_t == len(t)) || !found{
            return false
        }
    }   
    return true
}
