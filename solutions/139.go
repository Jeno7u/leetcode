package main

// идем снизу вверх. Если есть такое слово, что если его убрать остальная часть
// слова можно разбить, то вернем true. Но мы идем не от убирания слов, а к добавлении
func wordBreak(s string, wordDict []string) bool {
    dp := make([]bool, len(s) + 1)
    dp[len(s)] = true

    for i := len(s) - 1; i >= 0; i-- {
        for j := range wordDict {
            if (i + len(wordDict[j]) <= len(s)) && (s[i:i + len(wordDict[j])] == wordDict[j]) {
                dp[i] = dp[i + len(wordDict[j])]
            }
            if dp[i] == true {
                break
            }
        }
    }
    return dp[0]
}