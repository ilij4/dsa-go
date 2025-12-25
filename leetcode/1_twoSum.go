package leetcode

func twoSum(nums []int, target int) []int {
	remainMap := make(map[int]int)

	for i := range nums {
		remaining := target - nums[i]

		r, found := remainMap[remaining]

		if found {
			return []int{i, r}
		}

		remainMap[nums[i]] = i
	}

	return []int{}
}
