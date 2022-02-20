package weekgame

// https://leetcode-cn.com/contest/weekly-contest-281/problems/count-integers-with-even-digit-sum/

func CountEven(num int) (ans int) {
	for i := 1; i <= num; i++ {
		tmp, j := 0, i
		for j != 0 {
			tmp += j % 10
			j = j / 10
		}
		if tmp%2 == 0 {
			ans += 1
		}
	}
	return
}
