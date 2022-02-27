package weekgame

// https://leetcode-cn.com/contest/weekly-contest-282/problems/counting-words-with-a-given-prefix/

func prefixCount(words []string, pref string) (ans int) {
	for _, v := range words {
		lenW := len(v)
		lenP := len(pref)
		if lenW < lenP {
			continue
		}
		for i, j := 0, 0; j < lenP; j++ {
			if v[i] != pref[j] {
				break
			}
			if j == lenP-1 {
				ans += 1
			}
			i++
		}
	}
	return
}
