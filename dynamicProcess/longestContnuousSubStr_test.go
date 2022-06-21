package dynamicprocess

import (
	"fmt"
	"testing"
)

func TestLongestContnuousSubStr(t *testing.T) {
	strA, strB := "abcsleijjjjwoei", "bcewoejjjjk"
	comStr := longestContnuousSubStr(strA, strB)
	fmt.Println(comStr)
}
