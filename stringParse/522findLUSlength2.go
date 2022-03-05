package stringparse

// https://leetcode-cn.com/problems/longest-uncommon-subsequence-ii/

func findLUSlength2(strs []string) int {
	res := make([]map[string]int, 11)
	for i := range res {
		res[i] = make(map[string]int)
	}
	for _, s := range strs {
		n := len(s)
		res[n][s] += 1
	}
	for i := 10; i > 0; i-- {
		for k, v := range res[i] {
			if v == 1 {
				single := true
				for j := i + 1; j <= 10; j++ {
					for key, _ := range res[j] {
						if isSubstr(key, k) {
							single = false
							break
						}
					}
					if !single {
						break
					}
				}
				if single {
					return i
				}
			}
		}
	}
	return -1
}

func isSubstr(parent, child string) bool {
	lp, lc := len(parent), len(child)
	for i, j := 0, 0; i < lc; i++ {
		for j < lp {
			if parent[j] != child[i] {
				j++
			} else {
				break
			}
		}
		if j >= lp {
			return false
		}
		j++
	}
	return true
}
