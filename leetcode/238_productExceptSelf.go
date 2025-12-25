package leetcode

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	pre := 1
	result[0] = 1

	for i := 0; i < len(nums)-1; i++ {
		pre *= nums[i]
		result[i+1] = nums[i+1] * pre
	}

	post := 1
	for i := len(nums) - 1; i > 0; i-- {
		post *= nums[i]
		result[i-1] = nums[i-1] * post
	}

	return result

}
