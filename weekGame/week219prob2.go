package weekgame

// https://leetcode-cn.com/contest/weekly-contest-219/problems/partitioning-into-minimum-number-of-deci-binary-numbers/

func minPartitions(n string) int {
	ans := 0
	for _, ch := range n {
		if int(ch-'0') > ans {
			ans = int(ch - '0')
		}
	}
	return ans
}
