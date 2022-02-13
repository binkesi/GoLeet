package weekgame

// https://leetcode-cn.com/contest/weekly-contest-280/problems/count-operations-to-obtain-zero/

func countOperations(num1 int, num2 int) (ans int) {
	for num1 != 0 && num2 != 0 {
		if num1 >= num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
		ans += 1
	}
	return
}
