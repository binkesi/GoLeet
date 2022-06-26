package slidingwindow

// https://leetcode.cn/problems/substring-with-concatenation-of-all-words/

func findSubstring(s string, words []string) (ans []int) {
	wordL := len(words[0])
	wn := len(words)
	n := len(s)
	wMap := make(map[string]int)
	for _, w := range words {
		if _, ok := wMap[w]; !ok {
			wMap[w] = 1
		} else {
			wMap[w]++
		}
	}
	for i := 0; i <= n-wordL*wn; i++ {
		newMap := make(map[string]int)
		for k, v := range wMap {
			newMap[k] = v
		}
		for j := 0; j < wn; j++ {
			curWord := ""
			if i+(j+1)*wordL < n {
				curWord = s[i+j*wordL : i+(j+1)*wordL]
			} else {
				curWord = s[i+j*wordL:]
			}
			if v, ok := newMap[curWord]; ok {
				if v < 1 {
					break
				} else {
					newMap[curWord]--
				}
			} else {
				break
			}
			if j == wn-1 {
				ans = append(ans, i)
			}
		}
	}
	return
}
