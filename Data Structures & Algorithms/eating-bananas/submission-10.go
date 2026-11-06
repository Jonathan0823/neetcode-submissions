func minEatingSpeed(piles []int, h int) int {
	eatPerPile := 0
	for _, p := range piles { 
		eatPerPile = max(eatPerPile, p)
	}
	l, r := 1, eatPerPile
	for l <= r { 
		m := l + (r-l)/2
		time := eatBananas(piles, m)
		if time <= h {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func eatBananas(piles []int, speed int) int { 
	time := 0
	for _, p := range piles { 
		time += (p + speed - 1) / speed
	}

	return time
}
