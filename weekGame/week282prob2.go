package weekgame

// https://leetcode-cn.com/contest/weekly-contest-282/problems/minimum-number-of-steps-to-make-two-strings-anagram-ii/

func minSteps(s string, t string) (ans int) {
	mapA, mapB := map[rune]int{}, map[rune]int{}
	for _, b := range s {
		mapA[b] += 1
	}
	for _, b := range t {
		mapB[b] += 1
	}
	for k, v := range mapA {
		if _, ok := mapB[k]; !ok {
			ans += v
		} else if v > mapB[k] {
			ans += (v - mapB[k])
		}
	}
	for k, v := range mapB {
		if _, ok := mapA[k]; !ok {
			ans += v
		} else if v > mapA[k] {
			ans += (v - mapA[k])
		}
	}
	return
}
