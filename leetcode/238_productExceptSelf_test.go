package leetcode

import (
	"fmt"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	nums := []int{1,2,3,4}

	result := productExceptSelf(nums)

	fmt.Printf("%v", result)

	if len(result) != 4 {
		t.Errorf("Empty return")
	}

}