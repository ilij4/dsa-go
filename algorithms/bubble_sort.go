package algorithms

// BubbleSort returns a sorted copy of nums in ascending order using bubble sort.
func BubbleSort(nums []int) []int {
	sorted := append([]int(nil), nums...)
	if len(sorted) < 2 {
		return sorted
	}

	for i := 0; i < len(sorted)-1; i++ {
		swapped := false

		for j := 0; j < len(sorted)-1-i; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
				swapped = true
			}
		}

		if !swapped {
			return sorted
		}
	}

	return sorted
}
