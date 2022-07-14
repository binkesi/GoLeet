package trie

type WordFilter struct {
	pre *TrieNode
	suf *TrieNode
}

type TrieNode struct {
	ind []int
	arr [26]*TrieNode
}

func Constructor(words []string) WordFilter {
	pre := &TrieNode{ind: []int{}, arr: [26]*TrieNode{}}
	suf := &TrieNode{ind: []int{}, arr: [26]*TrieNode{}}
	for i, word := range words {
		build := pre
		for _, b := range word {
			num := b - 'a'
			if build.arr[num] == nil {
				build.arr[num] = &TrieNode{ind: []int{}, arr: [26]*TrieNode{}}
			}
			build = build.arr[num]
			build.ind = append(build.ind, i)
		}
	}
	for j, word := range words {
		build := suf
		l := len(word)
		for i := l - 1; i >= 0; i-- {
			num := rune(word[i]) - 'a'
			if build.arr[num] == nil {
				build.arr[num] = &TrieNode{ind: []int{}, arr: [26]*TrieNode{}}
			}
			build = build.arr[num]
			build.ind = append(build.ind, j)
		}
	}
	return WordFilter{pre: pre, suf: suf}
}

func (this *WordFilter) F(pref string, suff string) int {
	pre := this.pre
	suf := this.suf
	var preArr []int
	var sufArr []int
	for _, b := range pref {
		num := b - 'a'
		pre = pre.arr[num]
		if pre == nil {
			return -1
		}
	}
	preArr = pre.ind
	for i := len(suff) - 1; i >= 0; i-- {
		num := rune(suff[i]) - 'a'
		suf = suf.arr[num]
		if suf == nil {
			return -1
		}
	}
	sufArr = suf.ind
	for i, j := len(preArr)-1, len(sufArr)-1; i >= 0 && j >= 0; {
		if preArr[i] == sufArr[j] {
			return preArr[i]
		}
		if preArr[i] > sufArr[j] {
			i--
		} else {
			j--
		}
	}
	return -1
}
