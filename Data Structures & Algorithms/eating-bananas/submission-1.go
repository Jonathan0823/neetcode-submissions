func minEatingSpeed(piles []int, h int) int {
	maxSpeed := 0
	for _, p := range piles { 
		maxSpeed = max(maxSpeed, p)
	}

	l, r := 1, maxSpeed

	res := maxSpeed
	for l <= r { 
		m := l + (r-l)/2

		speedM := countEatTime(piles, m)
		if speedM <= h { 
			res = min(res, m)
			r = m - 1
		} else if speedM > h { 
			l = m + 1
		} 
	}

	return res

}

func countEatTime(piles []int, speed int) int { 
	res := 0
	for _, p := range piles { 
		res += (p + speed - 1)/ speed
	}

	return res
}
