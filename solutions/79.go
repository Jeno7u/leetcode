package main

// делаем просто dfs когда находим начало слова. Смотрим по соседям следующую букву
func exist(board [][]byte, word string) bool {
    for j := range board {
        for i := range board[0] {
            if board[j][i] == word[0] {
                found := checkWord(board, word, 0, i, j)
                if found {
                    return true
                }
            }
        }
    }
    return false
}

func checkWord(board [][]byte, word string, wordIdx int, i, j int) bool {
    if i < 0 || j < 0 || i >= len(board[0]) || j >= len(board) || board[j][i] != word[wordIdx] {
        return false
    }
    if wordIdx == len(word) - 1 {
        return true
    }

    tmp := board[j][i]
    board[j][i] = '*'

    found := checkWord(board, word, wordIdx + 1, i - 1, j) ||
             checkWord(board, word, wordIdx + 1, i, j - 1) ||
             checkWord(board, word, wordIdx + 1, i + 1, j) ||
             checkWord(board, word, wordIdx + 1, i, j + 1)

    board[j][i] = tmp

    return found
}