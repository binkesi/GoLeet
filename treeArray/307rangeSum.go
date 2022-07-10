package treearray

// https://leetcode.cn/problems/range-sum-query-mutable/submissions/

type NumArray struct {
	arr []int
	bits []int
}


func Constructor(nums []int) NumArray {
	n := len(nums)
	bits := make([]int, n+1)
	for i := 1; i <= n; i++{
		bits[i] = nums[i-1]
	}
	for j := 1; j <= n; j++{
		k := j + lowBits(j)
		if k <= n{
            bits[k] = bits[k] + bits[j]
        }
	}
	return NumArray{ arr: nums, bits: bits }
}

func lowBits(i int) int{
	return i & (-i)
}


func (this *NumArray) Update(index int, val int)  {
	diff := val - this.arr[index]
	this.arr[index] = val
	i := index + 1
	for i <= len(this.arr){
		this.bits[i] += diff
		i = i + lowBits(i)
	}
}


func (this *NumArray) SumRange(left int, right int) int {
	sum1, sum2 := 0, 0
	for i := left; i > 0;{
		sum1 += this.bits[i]
		i = i - lowBits(i)
	}
	for i := right+1; i > 0;{
		sum2 += this.bits[i]
		i = i - lowBits(i)
	}
	return sum2 - sum1
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */