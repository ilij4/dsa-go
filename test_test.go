package main

import (
	"fmt"
	"testing"
)

func Test_test(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2

	result := test(nums, k)

	fmt.Printf("%v", result)

	if result != 1 {
		t.Errorf("Wrong result")
	}

}
