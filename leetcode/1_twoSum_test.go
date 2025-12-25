package leetcode

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)

	fmt.Printf("%v", result)

	if len(result) != 2 {
		t.Errorf("Empty return")
	}

}
