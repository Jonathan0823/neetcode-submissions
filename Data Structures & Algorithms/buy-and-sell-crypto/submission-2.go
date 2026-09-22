func maxProfit(prices []int) int {
	l, maxProfit := 0, 0
	for r := 1; r < len(prices); r++ { 
		maxProfit = max(maxProfit, prices[r] - prices[l])

		if prices[r] < prices[l] { 
			l = r
		}
		
	} 

	return maxProfit

}
