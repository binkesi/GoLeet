package weekgame

// https://leetcode-cn.com/contest/weekly-contest-280/problems/minimum-operations-to-make-the-array-alternating/

func minimumOperations(nums []int) (ans int) {
	lenN := len(nums)
	if lenN == 1 {
		return 0
	}
	evenMap, oddMap := make(map[int]int), make(map[int]int)
	for i := 0; i < lenN; i++ {
		if i%2 == 0 {
			evenMap[nums[i]] += 1
		} else {
			oddMap[nums[i]] += 1
		}
	}
	eTime, oTime, eNum, oNum := 0, 0, 0, 0
	eTime1, oTime1, eNum1, oNum1 := 0, 0, 0, 0
	for k, v := range evenMap {
		if eTime <= v {
			eNum1 = eNum
			eTime1 = eTime
			eTime = v
			eNum = k
		} else if eTime1 <= v {
			eTime1 = v
			eNum1 = k
		}
	}
	for k, v := range oddMap {
		if oTime < v {
			oNum1 = oNum
			oTime1 = oTime
			oTime = v
			oNum = k
		} else if oTime1 <= v {
			oTime1 = v
			oNum1 = k
		}
	}
	if eNum == oNum {
		if eTime+oTime1 > eTime1+oTime {
			oNum = oNum1
			oTime = oTime1
		} else {
			eNum = eNum1
			eTime = eTime1
		}
	}
	ans = (lenN+1)/2 - eTime + lenN/2 - oTime
	return
}
