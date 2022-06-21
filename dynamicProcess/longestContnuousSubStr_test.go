package dynamicprocess

import (
	"fmt"
	"testing"
)

func TestLongestContnuousSubStr(t *testing.T) {
	strA, strB := "abcsleiwoei", "bcewoe"
	comStr := longestContnuousSubStr(strA, strB)
	fmt.Println(comStr)
}
