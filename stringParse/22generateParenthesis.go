package stringparse

// https://leetcode-cn.com/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	var res []string = make([]string, 0)
	putParenthesis("", n, 0, 0, &res)
	return res
}

func putParenthesis(curStr string, npairs, nleft, nright int, strs *[]string) {
	if nleft+nright == 2*npairs {
		*strs = append(*strs, curStr)
		return
	}
	if nright >= nleft {
		curStr += "("
		putParenthesis(curStr, npairs, nleft+1, nright, strs)
	} else if nleft == npairs {
		curStr += ")"
		putParenthesis(curStr, npairs, nleft, nright+1, strs)
	} else {
		tmp1 := curStr + "("
		putParenthesis(tmp1, npairs, nleft+1, nright, strs)
		tmp2 := curStr + ")"
		putParenthesis(tmp2, npairs, nleft, nright+1, strs)
	}
}
