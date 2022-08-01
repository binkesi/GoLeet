package dfsbfs

// https://leetcode.cn/problems/implement-magic-dictionary/submissions/

type MagicDictionary struct {
	root *trieNode
}

type trieNode struct {
	children map[byte]*trieNode
	isLeaf   bool
}

func Constructor() MagicDictionary {
	return MagicDictionary{root: &trieNode{children: make(map[byte]*trieNode)}}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, s := range dictionary {
		cur := this.root
		for i := range s {
			if _, ok := cur.children[s[i]]; !ok {
				cur.children[s[i]] = &trieNode{children: make(map[byte]*trieNode)}
			}
			cur = cur.children[s[i]]
		}
		cur.isLeaf = true
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	return dfs(this.root, searchWord, false)
}

func dfs(node *trieNode, word string, modified bool) bool {
	if word == "" {
		return modified && node.isLeaf
	}
	b := word[0]
	if node.children[b] != nil && dfs(node.children[b], word[1:], modified) {
		return true
	}
	if !modified {
		for k, v := range node.children {
			if k != b && v != nil && dfs(node.children[k], word[1:], true) {
				return true
			}
		}
	}
	return false
}
