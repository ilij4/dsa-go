package leetcode

import (
	"fmt"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	nums := []int{1,1,1,2,2,3}
	k := 2

	result := topKFrequent(nums, k)

	fmt.Printf("%v", result)

	if len(result) != 2 {
		t.Errorf("Empty return")
	}

}