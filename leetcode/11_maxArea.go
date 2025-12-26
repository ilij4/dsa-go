package leetcode

func maxArea(height []int) int {
	area := 0
	left, right := 0, len(height)-1
	for left < right {
		h := min(height[left], height[right])
		w := right - left
		currentArea := h * w
		if currentArea > area {
			area = currentArea
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area  
}
