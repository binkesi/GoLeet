package arraymanipulate

// https://leetcode-cn.com/problems/1-bit-and-2-bit-characters/

func isOneBitCharacter(bits []int) bool {
	lenB := len(bits)
	cnt := 0
	for i := lenB - 2; i >= 0; i-- {
		if bits[i] == 1 {
			cnt += 1
		} else {
			break
		}
	}
	return cnt%2 == 0
}
