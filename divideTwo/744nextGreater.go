package dividetwo

// https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target/

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	l, r := 0, n-1
	if target >= letters[r] {
		return letters[l]
	}
	for l < r {
		mid := l + (r-l)>>1
		if letters[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return letters[l]
}
