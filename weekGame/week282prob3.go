package weekgame

import "sort"

// https://leetcode-cn.com/contest/weekly-contest-282/problems/minimum-time-to-complete-trips/

func minimumTime(time []int, totalTrips int) (ans int64) {
	sort.Ints(time)
	lenT := len(time)
	l, r := time[0]*totalTrips/lenT, time[0]*totalTrips
	if l == r {
		ans = int64(l)
		return
	}
	for l < r {
		mid := l + (r-l)/2
		sum := 0
		for _, v := range time {
			if mid/v == 0 {
				break
			}
			sum += (mid / v)
		}
		if sum == totalTrips {
			ans = int64(mid)
			break
		}
		if sum < totalTrips {
			l = mid
			if l+1 == r {
				ans = int64(r)
				break
			}
		} else if sum > totalTrips {
			r = mid
			if r-1 == l {
				ans = int64(r)
				break
			}
		}
	}
	for k := int(ans); ; k-- {
		sum := 0
		for _, v := range time {
			if k/v == 0 {
				break
			}
			sum += (k / v)
		}
		if sum < totalTrips {
			ans = int64(k + 1)
			break
		}
	}
	return
}
