package leetcode

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	result := isPalindrome("A man, a plan, a canal: Panama")

	fmt.Printf("%v", result)

	if !result {
		t.Errorf("not palindrome")
	}
}

func Test_isPalindrome1(t *testing.T) {
	result := isPalindrome("0P")

	fmt.Printf("%v", result)

	if result {
		t.Errorf("not palindrome")
	}
}
