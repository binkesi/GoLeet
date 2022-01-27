package weekgame

// https://leetcode-cn.com/contest/weekly-contest-277/problems/maximum-good-people-based-on-statements/

func maximumGood(statements [][]int) (ans int) {
next:
	for i, n := 1, len(statements); i < 1<<n; i++ {
		cnts := 0
		for j := 0; j < n; j++ {
			if (i>>j)&1 != 0 {
				for k, s := range statements[j] {
					if s < 2 && (i>>k)&1 != s {
						continue next
					}
				}
				cnts++
			}
		}
		if cnts > ans {
			ans = cnts
		}
	}
	return
}
