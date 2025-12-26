package leetcode

import (
	"fmt"
	"testing"
)

func Test_ThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}

	result := threeSum(nums)

	fmt.Printf("%v", result)

	if len(result) != 2 {
		t.Errorf("Empty return")
	}

}
