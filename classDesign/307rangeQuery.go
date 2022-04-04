package classdesign

// https://leetcode-cn.com/problems/range-sum-query-mutable/

type NumArray struct {
	nums, tree []int
}

func ConstructorR(nums []int) NumArray {
	tree := make([]int, len(nums)+1)
	na := NumArray{nums, tree}
	for i, num := range nums {
		na.add(i+1, num)
	}
	return na
}

func (na *NumArray) add(index, val int) {
	for ; index < len(na.tree); index += index & -index {
		na.tree[index] += val
	}
}

func (na *NumArray) prefixSum(index int) (sum int) {
	for ; index > 0; index &= index - 1 {
		sum += na.tree[index]
	}
	return
}

func (na *NumArray) Update(index, val int) {
	na.add(index+1, val-na.nums[index])
	na.nums[index] = val
}

func (na *NumArray) SumRange(left, right int) int {
	return na.prefixSum(right+1) - na.prefixSum(left)
}
