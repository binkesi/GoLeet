package weekgame

// https://leetcode-cn.com/contest/weekly-contest-283/problems/cells-in-a-range-on-an-excel-sheet/

func cellsInRange(s string) (ans []string) {
	for i := s[0] - 'A'; i <= s[3]-'A'; i++ {
		for j := s[1] - '1'; j <= s[4]-'1'; j++ {
			tmp := []byte{}
			tmp = append(tmp, 'A'+i)
			tmp = append(tmp, '1'+j)
			ans = append(ans, string(tmp))
		}
	}
	return
}
