func maxArea(heights []int) int {
	maxArea, left := 0, 0
	right, width := len(heights) - 1, len(heights) - 1

	for left < right { 
		maxArea = max(maxArea, min(heights[left], heights[right]) * width)
		if heights[left]< heights[right] {
			 left++
		} else { 
			right--
		}
		width--
	}
	
	return maxArea
}
