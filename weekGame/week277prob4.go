package weekgame

// https://leetcode-cn.com/contest/weekly-contest-277/problems/maximum-good-people-based-on-statements/

func maximumGood(statements [][]int) int {
	lenP := len(statements[0])
	res := 0
	for i := 0; i < lenP; i++ {
		goodbad := make([]int, lenP)
		put(statements, i, &goodbad)
		tmp := 0
		for j := 0; j < len(goodbad); j++ {
			if goodbad[j] == 1 {
				tmp += 1
			}
		}
		if res < tmp {
			res = tmp
		}
	}
	return res
}

type pair struct{ idx, value int }

func put(state [][]int, ind int, result *[]int) {
	thisState := state[ind]
	lenS := len(thisState)
	tmp := []pair{{}}
	for i := 0; i < lenS; i++ {
		if thisState[i] == 0 {
			if (*result)[i] == 0 || (*result)[i] == 2 {
				(*result)[i] = 2
				tmp = append(tmp, pair{i, 2})
			} else {
				for _, v := range tmp {
					(*result)[v.idx] = v.value
				}
				return
			}
		} else if thisState[i] == 1 {
			if (*result)[i] == 0 || (*result)[i] == 1 {
				(*result)[i] = 1
				tmp = append(tmp, pair{i, 1})
			} else {
				for _, v := range tmp {
					(*result)[v.idx] = v.value
				}
				return
			}
		}
	}
	for k := 0; k < lenS; k++ {
		if (*result)[k] == 0 {
			put(state, k, result)
		}
	}
}
