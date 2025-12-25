package leetcode

func containsDuplicate(nums []int) bool {
	existingMap := make(map[int]bool, len(nums))

	for _, n := range nums {
		if _, ok := existingMap[n]; ok {
			return true
		}
		existingMap[n] = true
	}

	return false
}
