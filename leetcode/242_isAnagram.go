package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charMap := make(map[byte]uint, 26)

	for i := range len(s) {
		charMap[s[i]]++
		charMap[t[i]]--
	}

	for _, v := range charMap {
		if v != 0 {
			return false
		}
	}

	return true
}
