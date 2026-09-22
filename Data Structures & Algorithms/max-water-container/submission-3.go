func maxArea(heights []int) int {
	maxWater := 0

	l, r := 0, len(heights) - 1

	for l < r  {
		maxWater = max(maxWater, min(heights[l], heights[r]) * (r-l))
		if heights[l] < heights[r] { 
			l++
		} else { 
			r--
		}

	}

	return maxWater
}
