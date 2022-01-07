package goleet

//https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	digLen := len(digits)
	if digLen == 0 {
		return res
	}
	digMap := make(map[string][]string)
	digMap["2"] = []string{"a", "b", "c"}
	digMap["3"] = []string{"d", "e", "f"}
	digMap["4"] = []string{"g", "h", "i"}
	digMap["5"] = []string{"j", "k", "l"}
	digMap["6"] = []string{"m", "n", "o"}
	digMap["7"] = []string{"p", "q", "r", "s"}
	digMap["8"] = []string{"t", "u", "v"}
	digMap["9"] = []string{"w", "x", "y", "z"}
	for i := 0; i < len(digMap[string(digits[0])]); i++ {
		if digLen == 1 {
			res = append(res, digMap[string(digits[0])][i])
			continue
		}
	}
	return res
}
