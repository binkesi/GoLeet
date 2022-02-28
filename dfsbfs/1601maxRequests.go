package dfsbfs

// https://leetcode-cn.com/problems/maximum-number-of-achievable-transfer-requests/

func maximumRequests(n int, requests [][]int) (ans int) {
	lenR := len(requests)
	mask := (1 << lenR) - 1
	for i := 1; i <= mask; i++ {
		req := [][]int{}
		for k, tmp := lenR-1, i; tmp > 0; {
			if tmp&1 == 1 {
				req = append(req, requests[k])
			}
			k -= 1
			tmp = tmp >> 1
		}
		if isRing(req) && len(req) > ans {
			ans = len(req)
		}
	}
	return
}

func isRing(req [][]int) bool {
	mapR := map[int]int{}
	for _, v := range req {
		mapR[v[0]] += 1
		mapR[v[1]] -= 1
	}
	for _, v := range mapR {
		if v != 0 {
			return false
		}
	}
	return true
}
