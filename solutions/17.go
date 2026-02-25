package main


// стандартный бэ
func letterCombinations(digits string) []string {
    result := []string{}

    digitMap := map[byte][]byte{'2': []byte{'a', 'b', 'c'}, '3': []byte{'d', 'e', 'f'},
    '4': []byte{'g', 'h', 'i'}, '5': []byte{'j', 'k','l'}, '6': []byte{'m', 'n', 'o'}, 
    '7': []byte{'p', 'q', 'r', 's'}, '8': []byte{'t', 'u', 'v'}, '9': []byte{'w', 'x', 'y', 'z'}}
    curr := []byte{}
    var backtrack func(idx int)
    backtrack = func (idx int) {
        if idx == len(digits) {
            result = append(result, string(curr))
            return
        }
        for _, ch := range digitMap[digits[idx]] {
            curr = append(curr, ch)
            backtrack(idx + 1)
            curr = curr[:len(curr) - 1]
        }
    }

    backtrack(0)
    return result
}