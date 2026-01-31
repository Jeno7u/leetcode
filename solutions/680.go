package main


func validate(s string, l, r int) bool {
    for l < r {
        if s[l] != s[r] {
            return false
        }
        l, r = l + 1, r - 1
    }
    return true
}

func validPalindrome(s string) bool {
    l, r := 0, len(s) - 1
    for l < r {
        if s[l] != s[r] {
            return validate(s, l+1, r) || validate(s, l, r - 1)
        }
        l, r = l + 1, r - 1
    }
    return true
}