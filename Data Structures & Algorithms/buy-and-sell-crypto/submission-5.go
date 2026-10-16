func maxProfit(prices []int) int {
	l, res := 0, 0
	for idx, r := range prices { 
		res = max(res, r - prices[l])
		if r < prices[l] {
			l= idx
		}
	}

	return res
}
