package leetcode

func groupAnagrams(strs []string) [][]string {

	groups := make(map[[26]uint16][]string)

	for _, str := range strs {
		var hash [26]uint16
		for _, s := range str {
			hash[s-'a']++
		}

		if _, ok := groups[hash]; ok {
			groups[hash] = append(groups[hash], str)
		} else {
			groups[hash] = []string{str}
		}
	}

	var result [][]string

	for _, v := range groups {
		result = append(result, v)
	}

	return result
}
