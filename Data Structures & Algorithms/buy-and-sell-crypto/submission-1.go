func maxProfit(prices []int) int {
	l, r := 0, 1
	profit := 0
	for r < len(prices) { 
		profit = max(profit, prices[r] - prices[l])
		if prices[r] < prices[l] { 
			l = r
		}
		r++
	}

	return profit

}
