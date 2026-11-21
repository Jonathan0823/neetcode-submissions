func maxProfit(prices []int) int {
	res := 0
	l := 0
	for r := 1; r < len(prices); r++ { 
		if prices[l] > prices[r] { 
			l = r
			continue
		}

		res = max(res, prices[r] - prices[l])
	}

	return res
}
