package leetcode

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int, len(nums))

	for _, num := range nums {
		freq[num]++
	}

	groups := make([][]int, len(nums))
	for n, f := range freq {
		groups[f] = append(groups[f], n)
	}

    res := make([]int, 0, k)
	for f := len(groups) - 1; f >= 1 && len(res) < k; f-- {
		for _, num := range groups[f] {
			res = append(res, num)
			if len(res) == k {
				break
			}
		}
	}
	return res

}
