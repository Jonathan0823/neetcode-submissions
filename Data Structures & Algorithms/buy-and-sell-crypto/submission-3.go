func maxProfit(prices []int) int {
	res := 0
	l := 0
	for idx, price := range prices { 
		res = max(res, price - prices[l])
		if price < prices[l] { 
			l = idx
		}
	} 

	return res
}
