package main


type TrieNode struct {
    children [26]*TrieNode
    endOfWord bool
}

func (node *TrieNode) Insert(word string) {
    currNode := node
    for i := range word {
        idx := word[i] - 'a'
        if currNode.children[idx] == nil {
            currNode.children[idx] = &TrieNode{}
        }
        currNode = currNode.children[idx]
    }
    currNode.endOfWord = true
}

// Составляем префиксное дерево и заполняем его словами. Потом мы делаем dfs поиск для каждой буквы внутри board. 
// По сути на каждом этапе dfs мы имеем слово которое мы собрали на текущий момент и мы можем проверить через префиксное дерево
// является ли оно префиксом. Конечно вместо того чтобы делать поиск для одних и тех же префиксов на рекурсивных вызовах мы 
// можем делать node = node.children[char] и проверять только является ли node.endOfWord == true. Если да, то добавляем 
// его в result (то есть нашли как раз) и ставим node.endOfWord = false чтобы не добавлять повторные слова
func findWords(board [][]byte, words []string) []string {
    trieRoot := TrieNode{}
    for i := range words {
        trieRoot.Insert(words[i])
    }

    result := []string{}
    var dfs func(i, j int, trie *TrieNode, currWord []byte)
    dfs = func(i, j int, trie *TrieNode, currWord []byte) {
        if i < 0 || j < 0 || i == len(board[0]) || j == len(board) || board[j][i] == '*' || trie.children[board[j][i] - 'a'] == nil {
            return
        }

        tmp := board[j][i]
        board[j][i] = '*'

        currWord = append(currWord, tmp)
        trie = trie.children[tmp - 'a']
        if trie.endOfWord {
			result = append(result, string(currWord))
            trie.endOfWord = false
        }

        dfs(i - 1, j, trie, currWord)
        dfs(i, j + 1, trie, currWord)
        dfs(i + 1, j, trie, currWord)
        dfs(i, j - 1, trie, currWord)
    
        board[j][i] = tmp
    }

    for j := range board {
        for i := range board[0] {
            dfs(i, j, &trieRoot, []byte{})
        }
    }

    return result
}