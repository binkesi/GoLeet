package weekgame

// https://leetcode.cn/contest/weekly-contest-301/problems/smallest-number-in-infinite-set/

type SmallestInfiniteSet struct {
	nums []int
}

func ConstructorInf() SmallestInfiniteSet {
	res := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		res[i] = i + 1
	}
	return SmallestInfiniteSet{nums: res}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	for i := 0; i < 1000; i++ {
		if this.nums[i] != 0 {
			ans := this.nums[i]
			this.nums[i] = 0
			return ans
		}
	}
	return 0
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	for i := 0; i < 1000; i++ {
		if this.nums[i] == num {
			break
		}
		if i == num {
			this.nums[i-1] = num
			break
		}
	}
}
