package main

func localBigest(l, r, currentMax int, s string) string {
    for l >= 0 && r < len(s) && s[l] == s[r] {
        l--
        r++
    }
    if (r - l - 1) > currentMax {
        return s[l + 1:r]
    }
    return ""
}

// пробегаемся по строке и каждый индекс считаем центром палиндрома. Делаем обычную проверку на палиндром (localBigest).
// Когда s[l:r + 1] не является палиндромом, мы сужаем границы на один, то есть на предыдущую итерацию, где оно все еще палиндром.
// Также приходится делать обход два раза. Для нечетного кол-ва символов в палиндроме и для четного (центр 1 или 2 символа).
func longestPalindrome(s string) string {
    result := string(s[0])

	for i := 0; i < len(s); i++ {
		l, r := i - 1, i + 1
        local := localBigest(l, r, len(result), s)
        if local != "" {
            result = local
        }

        l, r = i, i + 1
		local = localBigest(l, r, len(result), s)
        if local != "" {
            result = local
        }
	}

    return result
}
