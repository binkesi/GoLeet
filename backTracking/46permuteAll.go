package backtracking

// https://leetcode-cn.com/problems/permutations/

func permute(nums []int) [][]int {
	var res [][]int = make([][]int, 0)
	putNums(nums, make([]int, 0), &res)
	return res
}

func putNums(nums, curArr []int, res *[][]int) {
	lenN, lenC := len(nums), len(curArr)
	if lenC == lenN {
		*res = append(*res, curArr)
		return
	}
	leftNum := make([]int, 0)
	if lenC == 0 {
		leftNum = nums
	}
	for _, v := range nums {
		for i, val := range curArr {
			if val == v {
				break
			}
			if i == lenC-1 {
				leftNum = append(leftNum, v)
			}
		}
	}
	for _, v := range leftNum {
		tmp := append(curArr, v)
		putNums(nums, tmp, res)
	}
}

func permute2(nums []int) (ans [][]int) {
	n := len(nums)
	var chooseNum func(ch []bool, cur []int)
	chooseNum = func(ch []bool, cur []int) {
		l := len(cur)
		if l == n {
			ans = append(ans, cur)
			return
		}
		for i, v := range ch {
			if !v {
				nch := make([]bool, n)
				ncur := make([]int, l)
				copy(nch, ch)
				copy(ncur, cur)
				nch[i] = true
				ncur = append(ncur, nums[i])
				chooseNum(nch, ncur)
			}
		}
	}
	for i := range nums {
		choice := make([]bool, n)
		current := make([]int, 0)
		choice[i] = true
		current = append(current, nums[i])
		chooseNum(choice, current)
	}
	return
}
