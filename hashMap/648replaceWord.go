package hashmap

import "strings"

// https://leetcode.cn/problems/replace-words/

type trieNode struct {
	val      int
	children map[int]*trieNode
	isLeaf   bool
}

func replaceWords(dictionary []string, sentence string) string {
	root := trieNode{val: -1, children: make(map[int]*trieNode)}
	for _, w := range dictionary {
		tmp := root
		l := len(w)
		for i := 0; i < l; i++ {
			ind := int(w[i] - 'a')
			if _, ok := tmp.children[ind]; ok {
				if i == l-1 {
					tmp.children[ind].isLeaf = true
				}
				tmp = *tmp.children[ind]
				continue
			}
			if i != l-1 {
				child := trieNode{val: ind, children: make(map[int]*trieNode)}
				tmp.children[ind] = &child
				tmp = child
			} else {
				child := trieNode{val: ind, children: make(map[int]*trieNode), isLeaf: true}
				tmp.children[ind] = &child
			}
		}
	}
	words := strings.Split(sentence, " ")
	for i, word := range words {
		tmp := root
		l := len(word)
		for j := range word {
			ind := int(word[j] - 'a')
			if _, ok := tmp.children[ind]; !ok {
				break
			} else {
				if tmp.children[ind].isLeaf == true {
					if j != l-1 {
						words[i] = words[i][0 : j+1]
					}
					break
				}
			}
			tmp = *tmp.children[ind]
		}
	}
	ans := strings.Join(words, " ")
	return ans
}
