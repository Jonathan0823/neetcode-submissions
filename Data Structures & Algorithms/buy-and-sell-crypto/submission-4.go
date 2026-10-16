func maxProfit(prices []int) int {
	l, res := 0, 0
	for r := 1; r < len(prices); r++ { 
		res = max(res, prices[r] - prices[l])
		if prices[r] < prices[l] {
			l=r
		}
	}

	return res
}
