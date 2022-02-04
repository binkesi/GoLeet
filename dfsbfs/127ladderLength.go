package dfsbfs

// https://leetcode-cn.com/problems/word-ladder/

func ladderLength(beginWord string, endWord string, wordList []string) (ans int) {
	for i, v := range wordList {
		if v == endWord {
			break
		}
		if i == len(wordList)-1 {
			return 0
		}
	}
	q1, q2 := []string{beginWord}, []string{endWord}
	h1, h2 := make(map[string]int), make(map[string]int)
	h1[beginWord] = 0
	h2[endWord] = 0
	for len(q1) != 0 && len(q2) != 0 {
		p1, p2 := q1[0], q2[0]
		l1, l2 := len(q1), len(q2)
		if l1 > l2 {
			q2 = q2[1:]
			for _, v := range wordList {
				if _, ok := h2[v]; !ok && canConvert(p2, v) {
					h2[v] = h2[p2] + 1
					if _, ok := h1[v]; ok {
						ans = h1[v] + h2[v] + 1
						return
					}
					q2 = append(q2, v)
				}
			}
		} else {
			q1 = q1[1:]
			for _, v := range wordList {
				if _, ok := h1[v]; !ok && canConvert(p1, v) {
					h1[v] = h1[p1] + 1
					if _, ok := h2[v]; ok {
						ans = h1[v] + h2[v] + 1
						return
					}
					q1 = append(q1, v)
				}
			}
		}
	}
	return 0
}

func canConvert(a, b string) bool {
	cnt := 0
	lenA := len(a)
	for i := 0; i < lenA; i++ {
		if a[i] != b[i] {
			cnt += 1
		}
	}
	return cnt == 1
}
