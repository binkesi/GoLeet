package weekgame

// https://leetcode-cn.com/contest/biweekly-contest-72/problems/find-three-consecutive-integers-that-sum-to-a-given-number/

func sumOfThree(num int64) (ans []int64) {
	if num%3 != 0 {
		return
	} else {
		res := num / 3
		ans = append(ans, res-1, res, res+1)
	}
	return
}
