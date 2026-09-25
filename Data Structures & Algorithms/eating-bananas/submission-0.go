func minEatingSpeed(piles []int, h int) int {
	maxPiles := 0
	for _, pile := range piles { 
		maxPiles = max(maxPiles, pile)
	}

	l, r := 1, maxPiles

	ans := maxPiles
	for l <= r { 
		mid := l + (r-l)/2
		totalHour := 0
		for _, pile := range piles { 
			totalHour += (pile + mid - 1) / mid
		}
		if totalHour > h { 
			l = mid + 1
		} else { 
			ans = min(ans, mid)
			r = mid -1
		}
	}

	return ans

}
