package arraymanipulate

// https://leetcode-cn.com/problems/corporate-flight-bookings/

func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n+1)
	var add func(int, int, int)
	add = func(i1, i2, i3 int) {
		res[i1-1] += i3
		res[i2] -= i3
	}
	for _, v := range bookings {
		add(v[0], v[1], v[2])
	}
	for i := 1; i < n; i++ {
		res[i] += res[i-1]
	}
	return res[0:n]
}
