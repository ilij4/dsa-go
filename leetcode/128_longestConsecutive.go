package leetcode

func longestConsecutive(nums []int) int {
	numSet := make(map[int]struct{})

	maxCount := 0

	for _, n := range nums {
		numSet[n] = struct{}{}
	}

	for n := range numSet {
		if _, ok := numSet[n-1]; ok {
			continue
		}

		count := 1
		for {
			if _, ok := numSet[n+count]; ok {
				count++
			} else {
				if count > maxCount {
					maxCount = count
				}
				break
			}
		}
	}

    return maxCount

}

func longestConsecutive1(nums []int) int {
	numSet := make(map[int]struct{}, len(nums))

	longest := 0

	for _, n := range nums {
		numSet[n] = struct{}{}
	}

	for x := range numSet {
		if _, hasPrev := numSet[x-1]; hasPrev {
			continue
		}

		length := 1
		for curr := x + 1; ; curr++ {
			if _, ok := numSet[curr]; ok {
				length++
			} else {
				break
			}
		}
		if length > longest {
			longest = length
		}
	}

    return longest
}