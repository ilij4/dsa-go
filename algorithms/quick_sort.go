package algorithms

func QuickSort(nums []int) []int {
	sorted := append([]int(nil), nums...)
	if len(sorted) < 2 {
		return sorted
	}

	quickSort(sorted, 0, len(sorted)-1)
	return sorted
}

func quickSort(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}

	pivot := nums[lo+(hi-lo)/2]
	i, j := lo, hi

	for i <= j {
		for nums[i] < pivot {
			i++
		}
		for nums[j] > pivot {
			j--
		}
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	if lo < j {
		quickSort(nums, lo, j)
	}
	if i < hi {
		quickSort(nums, i, hi)
	}
}
