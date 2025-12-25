package algorithms

func MergeSort(nums []int) []int {
	sorted := append([]int(nil), nums...)
	if len(sorted) < 2 {
		return sorted
	}

	buf := make([]int, len(sorted))
	mergeSort(sorted, buf, 0, len(sorted))
	return sorted
}

func mergeSort(nums, buf []int, lo, hi int) {
	if hi-lo <= 1 {
		return
	}

	mid := lo + (hi-lo)/2
	mergeSort(nums, buf, lo, mid)
	mergeSort(nums, buf, mid, hi)

	i, j, k := lo, mid, lo
	for i < mid && j < hi {
		if nums[i] <= nums[j] {
			buf[k] = nums[i]
			i++
		} else {
			buf[k] = nums[j]
			j++
		}
		k++
	}

	for i < mid {
		buf[k] = nums[i]
		i++
		k++
	}

	for j < hi {
		buf[k] = nums[j]
		j++
		k++
	}

	copy(nums[lo:hi], buf[lo:hi])
}
