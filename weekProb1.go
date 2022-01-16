package goleet

func divideString(s string, k int, fill byte) []string {
	lenS := len(s)
	sLoop, sLeft := lenS/k, lenS%k
	var res []string = make([]string, 0)
	for i := 0; i < sLoop; i++ {
		res = append(res, s[i*k:i*k+k])
	}
	if sLeft != 0 {
		fillLen := k - sLeft
		var fillBytes []byte = make([]byte, 0)
		for i := 0; i < fillLen; i++ {
			fillBytes = append(fillBytes, fill)
		}
		leftString := s[sLoop*k:] + string(fillBytes)
		res = append(res, leftString)
	}
	return res
}
