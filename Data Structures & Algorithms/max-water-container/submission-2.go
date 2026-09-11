func maxArea(heights []int) int {
	l, r, maxWater := 0, len(heights) - 1, 0

	for l < r { 
		maxWater = max(maxWater,  min(heights[l], heights[r]) * (r - l))
		
		if heights[l] < heights[r] { 
			l++
		} else { 
			r--
		}
	}

	return maxWater

}
