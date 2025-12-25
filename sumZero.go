package main

import (
	"fmt"
	"strings"
)

func sumZero(num1, num2, num3 float64) bool {

	str1 := fmt.Sprintf("%f", num1)
	str2 := fmt.Sprintf("%f", num2)
	str3 := fmt.Sprintf("%f", num3)

	parts1 := strings.Split(str1, ".")
	parts2 := strings.Split(str2, ".")
	parts3 := strings.Split(str3, ".")

	maxPrec := 1

	if len(parts1) > 1 && len(parts1[1]) > maxPrec {
		maxPrec = len(parts1[1])
	}

	if len(parts2) > 1 && len(parts2[1]) > maxPrec {
		maxPrec = len(parts2[1])
	}

	if len(parts3) > 1 && len(parts3[1]) > maxPrec {
		maxPrec = len(parts3[1])
	}

	multiplier := 10 ^ maxPrec

	sum := int64(num1 * float64(multiplier))
	sum1 := sum + int64(num2*float64(multiplier))
	sum2 := sum1 + int64(num3*float64(multiplier))

	if sum2 == 0 {
		return true
	}

	return false
}
