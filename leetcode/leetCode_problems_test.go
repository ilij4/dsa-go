package leetcode

import (
	"fmt"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}

	result := maxProfit(prices)

	fmt.Printf("%v", result)

	if result != 5 {
		t.Errorf("Wrong return")
	}
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	// s := "tmmzuxt"
	s := "abcabcbb"

	result := lengthOfLongestSubstring(s)

	fmt.Printf("%v", result)

	if result != 5 {
		t.Errorf("Wrong return")
	}
}

func Test_characterReplacement(t *testing.T) {
	s := "AABABBA"
	k := 1

	result := characterReplacement(s, k)

	fmt.Printf("%v", result)

	if result != 4 {
		t.Errorf("Wrong return")
	}
}
func Test_checkInclusion(t *testing.T) {
	s1 := "adc"
	// s2 := "eidbaooo"
	s2 := "dcda"

	result := checkInclusion(s1, s2)

	fmt.Printf("%v", result)

	if !result {
		t.Errorf("Wrong return")
	}
}

func Test_isValid(t *testing.T) {
	// s1 := "([])"
	s1 := "["

	result := isValid(s1)

	fmt.Printf("%v", result)

	if !result {
		t.Errorf("Wrong return")
	}
}
