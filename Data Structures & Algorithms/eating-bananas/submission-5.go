func minEatingSpeed(piles []int, h int) int {
	res := 0
	for _, p := range piles { 
		res = max(res, p)
	}

	l, r := 1, res
	for l <= r { 
		m := l + (r-l)/2
		speed := countTime(piles, m)
		if speed <= h { 
			res = min(res, m)
			r = m - 1
		} else { 
			l = m + 1
		}
	}

	return res
}

func countTime(piles []int, speed int) int { 
	time := 0
	for _, p := range piles { 
		time += (p + speed - 1)/ speed
	}

	return time
}
