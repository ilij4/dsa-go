package leetcode

import (
	"strings"
)

func isPalindrome(s string) bool {
	newS := strings.ToLower(s)

	i := 0
	j := len(newS) - 1

	for i < j {

		if (newS[i] < 'a' || newS[i] > 'z') && (newS[i] < '0' || newS[i] > '9') {
			i++
			continue
		}

		if (newS[j] < 'a' || newS[j] > 'z') && (newS[j] < '0' || newS[j] > '9') {
			j--
			continue
		}

		if newS[i] != newS[j] {
			return false
		}
		i++
		j--
	}

	return true
}
