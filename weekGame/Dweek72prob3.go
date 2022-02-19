package weekgame

// https://leetcode-cn.com/contest/biweekly-contest-72/problems/maximum-split-of-positive-even-integers/

func maximumEvenSplit(finalSum int64) (ans []int64) {
	if finalSum%2 != 0 {
		return
	} else {
		for k := int64(2); finalSum > 2*int64(k); k += 2 {
			ans = append(ans, k)
			finalSum -= k
		}
		ans = append(ans, finalSum)
	}
	return
}
