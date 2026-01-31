package main

func generateParenthesis(n int) []string {
    stack := make([]byte, 0, n * 2)
    result := []string{}

    var backtrack func(openN, closedN int)
    backtrack = func(openN, closedN int) {
        if n == openN && n == closedN {
            result = append(result, string(stack))
            return
        }

        if openN < n {
            stack = append(stack, '(')
            backtrack(openN + 1, closedN)
            stack = stack[:len(stack) - 1]
        }

        if closedN < openN {
            stack = append(stack, ')')
            backtrack(openN, closedN + 1)
            stack = stack[:len(stack) - 1]
        }
    }

    backtrack(0, 0)
    return result
}