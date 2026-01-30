package main

// status (-1 (less), 0(equal), 1(too much)) + which char
func checkArrays(arr1, arr2, curr1, curr2 [26]int) (int, byte) {
	for i := 0; i < 26; i++ {
		if arr1[i] != curr1[i] {
			if arr1[i] < curr1[i] {
				return 1, byte(i + int('a'))
			} else {
				return -1, byte(i + int('a'))
			}
		}

		if arr2[i] != curr2[i] {
			if arr2[i] < curr2[i] {
				return 1, byte(i + int('a'))
			} else {
				return -1, byte(i + int('a'))
			}
		}
	}
	return 0, byte(0)
}

func addValue(arr [26]int, value byte) {
	if 
}

func minWindow(s string, t string) string {
    if len(s) < len(t) {
        return ""
    }

    charCountLower := [26]int{}
    charCountUpper := [26]int{}

    for i := 0; i < len(t); i++ {
        if int(t[i]) < int('Z') {
            charCountUpper[int(t[i]) - int('A')]++
        } else {
            charCountLower[int(t[i]) - int('a')]++
        }
    }

    charCountLowerCurr := [26]int{}
    charCountUpperCurr := [26]int{}

    for i := 0; i < len(t); i++ {
        if int(s[i]) < int('Z') {
            charCountUpperCurr[int(s[i]) - int('A')]++
        } else {
            charCountLowerCurr[int(s[i]) - int('a')]++
        }
    }


    l, r := 0, len(t) - 1
    result := [2]int{}
    for r < len(s) {
		status, value := checkArrays(charCountLower, charCountUpper, charCountLowerCurr, charCountUpperCurr)
        if status == 0 {
            if result[1] - result[0] < r - l {
                result[0], result[1] = l, r
            }
        } else if status == -1 {
			for r < len(s) && s[r] != value {
				r++
			}
			if r == len(s) {
				break
			}
			
		} else if status == 1 {
			for l < len(s) && s[l] != value {
				l++
			}
		}
    }

	return s[result[0]:result[1] + 1]
}