func maxProfit(prices []int) int {
	left := 0;
	right := 1;

	maxProfit := 0;

	for right < len(prices) { 
		profit := prices[right] - prices[left];
		if profit > maxProfit {
			maxProfit = profit 
		}
		if (prices[left] > prices[right]) { 
			left++
		} else { 
			right++
		}
		
	}
	return maxProfit
}
