package leetcode

import (
	"github.com/ilij4/go-demo/helpers"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 121. Best Time to Buy and Sell Stock
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxProfit := 0
	i := 0
	j := 1

	for i < j && j < len(prices) {
		if prices[j] < prices[i] {
			i = j
		} else {
			profit := prices[j] - prices[i]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
		j++
	}

	return maxProfit
}

// 3. Longest Substring Without Repeating Characters
func lengthOfLongestSubstring(s string) int {
	longest := 0
	lastIndex := make(map[byte]int)

	start := 0

	for end := 0; end < len(s); end++ {
		if _, ok := lastIndex[s[end]]; ok {
			start = max(lastIndex[s[end]]+1, start)
		}

		longest = max(end-start+1, longest)

		lastIndex[s[end]] = end
	}

	return longest
}

// 424. Longest Repeating Character Replacement
func characterReplacement(s string, k int) int {
	var windowStart, maxLength, maxWindowCharCount int

	freqMap := make(map[byte]int)

	for windowEnd := 0; windowEnd < len(s); windowEnd++ {
		charRight := s[windowEnd]
		freqMap[charRight] = freqMap[charRight] + 1
		maxWindowCharCount = max(maxWindowCharCount, freqMap[charRight])

		currWindowLength := windowEnd - windowStart + 1
		for currWindowLength-maxWindowCharCount > k {
			charLeft := s[windowStart]
			freqMap[charLeft] = freqMap[charLeft] - 1
			windowStart++
			currWindowLength = windowEnd - windowStart + 1
		}
		maxLength = max(maxLength, currWindowLength)
	}

	return maxLength
}

// 567. Permutation in String
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Chars := make([]int, 26)
	s2Chars := make([]int, 26)

	s1Len := len(s1)

	for i := range s1Len {
		s1Chars[s1[i]-'a']++
		s2Chars[s2[i]-'a']++
	}

	matches := 0
	for i := range s1Chars {
		if s1Chars[i] == s2Chars[i] {
			matches++
		}
	}

	for windowEnd := s1Len; windowEnd < len(s2); windowEnd++ {
		if matches == 26 {
			return true
		}

		charRight := s2[windowEnd] - 'a'
		s2Chars[charRight]++
		if s2Chars[charRight] == s1Chars[charRight] {
			matches++
		} else if s1Chars[charRight]+1 == s2Chars[charRight] {
			matches--
		}

		charLeft := s2[windowEnd-s1Len] - 'a'
		s2Chars[charLeft]--
		if s2Chars[charLeft] == s1Chars[charLeft] {
			matches++
		} else if s1Chars[charLeft]-1 == s2Chars[charLeft] {
			matches--
		}
	}

	if matches == 26 {
		return true
	}

	return false
}

// 20. Valid Parentheses
func isValid(s string) bool {
	stack := helpers.NewStack[byte]()

	for i := 0; i < len(s); i++ {
		c := s[i]

		switch c {
		case '(':
			stack.Push(c)
		case '[':
			stack.Push(c)
		case '{':
			stack.Push(c)
		case ')':
			last, ok := stack.Pop()

			if !ok || last != '(' {
				return false
			}
		case ']':
			last, ok := stack.Pop()
			if !ok || last != '[' {
				return false
			}
		case '}':
			last, ok := stack.Pop()
			if !ok || last != '{' {
				return false
			}
		default:
			return false
		}
	}

	return stack.IsEmpty()
}
