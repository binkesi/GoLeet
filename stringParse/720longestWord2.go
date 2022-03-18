package stringparse

// https://leetcode-cn.com/problems/longest-word-in-dictionary/

func longestWord2(words []string) (ans string) {
	trie := Trie{}
	for _, word := range words {
		trie.Insert(word)
	}
	for _, word := range words {
		if trie.Search(word) {
			if len(word) > len(ans) {
				ans = word
			} else if len(word) == len(ans) && word < ans {
				ans = word
			}
		}
	}
	return
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func (t *Trie) Insert(word string) {
	node := t
	for _, w := range word {
		ch := w - 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, w := range word {
		ch := w - 'a'
		if node.children[ch] == nil || node.children[ch].isEnd == false {
			return false
		} else {
			node = node.children[ch]
		}
	}
	return true
}
