package main

type AnotherTrie struct {
	children  [26]*AnotherTrie
	endOfWord bool
}

func ConstructorTrie() AnotherTrie {
	return AnotherTrie{}
}

func (this *AnotherTrie) Insert(word string) {
	curr := this
	for i := range word {
		idx := word[i] - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &AnotherTrie{}
		}
		curr = curr.children[idx]
	}
	curr.endOfWord = true
}

func (this *AnotherTrie) Search(word string) bool {
	curr := this
	for i := range word {
		idx := word[i] - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return curr.endOfWord
}

func (this *AnotherTrie) StartsWith(prefix string) bool {
	curr := this
	for i := range prefix {
		idx := prefix[i] - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return true
}

/**
 * Your AnotherTrie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
