package main

import (
	"fmt"
	"testing"
)

func Test_sumZero(t *testing.T) {
	res := sumZero(1, 2, 3)
	fmt.Printf("%v", res)

	if res {
		t.Errorf("Not zero")
	}
}

func Test_sumZero1(t *testing.T) {
	res := sumZero(1.0, -2.99999, 3.000000001)
	fmt.Printf("%v", res)
}

func Test_sumZero2(t *testing.T) {
	res := sumZero(1.0, -2, 1)
	fmt.Printf("%v", res)
}

func Test_sumZero3(t *testing.T) {
	res := sumZero(0.000000001, 0.000000002, -0.000000003)
	fmt.Printf("%v", res)
}
