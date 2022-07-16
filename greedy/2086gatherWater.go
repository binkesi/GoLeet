package greedy

// https://leetcode.cn/problems/minimum-number-of-buckets-required-to-collect-rainwater-from-houses/submissions/

func minimumBuckets(street string) (ans int) {
	l := len(street)
	buckets := []byte(street)
	if l == 1 {
		if buckets[0] == 'H' {
			return -1
		} else {
			return 0
		}
	}
	for i := 0; i < l; i++ {
		if buckets[i] == 'H' {
			if i == 0 {
				if buckets[i+1] == 'H' {
					return -1
				} else {
					buckets[i+1] = 'B'
					ans++
				}
			} else if i == l-1 {
				if buckets[i-1] == 'H' {
					return -1
				}
				if buckets[i-1] == '.' {
					buckets[i-1] = 'B'
					ans++
				}
			} else {
				if buckets[i-1] == 'B' {
					continue
				}
				if buckets[i-1] == 'H' && buckets[i+1] == 'H' {
					return -1
				}
				if buckets[i-1] == '.' && buckets[i+1] == 'H' {
					buckets[i-1] = 'B'
					ans++
				} else {
					buckets[i+1] = 'B'
					ans++
				}
			}
		}
	}
	return
}
