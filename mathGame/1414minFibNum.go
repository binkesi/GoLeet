package mathgame

// https://leetcode-cn.com/problems/find-the-minimum-number-of-fibonacci-numbers-whose-sum-is-k/

func findMinFibonacciNumbers(k int) int {
	fib := []int{1, 1}
	ans := 1
	for i, sum := 0, 2; sum <= k; i++ {
		if sum == k {
			return 1
		}
		fib = append(fib, sum)
		sum = fib[i+1] + fib[i+2]
	}
	lenF := len(fib)
	ind, tmp := lenF-1, k-fib[lenF-1]
	for tmp != 0 {
		ans += 1
		i := findfib(fib[0:ind], tmp)
		ind = i
		tmp = tmp - fib[i]
	}
	return ans
}

func findfib(fib []int, k int) int {
	lenF := len(fib)
	if lenF <= 2 {
		return 0
	}
	left, right := 1, lenF-1
	mid := left + (right-left)/2
	for left < right {
		if fib[mid] == k {
			return mid
		}
		if fib[mid-1] == k {
			return mid - 1
		}
		if fib[mid+1] == k {
			return mid + 1
		}
		if fib[mid] > k && fib[mid-1] < k {
			return mid - 1
		} else if fib[mid] < k && fib[mid+1] > k {
			return mid
		} else if fib[mid] > k && fib[mid-1] > k {
			right = mid - 1
			mid = left + (right-left)/2
		} else if fib[mid] < k && fib[mid+1] < k {
			left = mid + 1
			mid = left + (right-left)/2
		}
	}
	return mid
}
