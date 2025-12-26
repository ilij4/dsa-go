package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		findTriplets(nums, i, &result)
	}

	return result
}

func findTriplets(numbers []int, firstInd int, result *[][]int) {
	left, right := firstInd+1, len(numbers)-1

	target := -numbers[firstInd]

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			*result = append(*result, []int{numbers[firstInd], numbers[left], numbers[right]})
            left++
            right--
            
            for left < right && numbers[left] == numbers[left-1] {
                left++
            }
            for left < right && numbers[right] == numbers[right+1] {
                right--
            }
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
}
