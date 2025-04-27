package arraystring

func maxProfit2(prices []int) int {
	//actions allowed:
	/*
	   1. buy
	   2. sell
	   3. only one share of stock
	   4. buy and sell on the same day
	*/

	if len(prices) <= 1 {
		return 0
	}
	n := len(prices)
	profit := 0
	cost := prices[0]
	for i := 1; i < n; i++ {
		if prices[i]-cost > 0 {
			profit += prices[i] - cost
		}
		cost = prices[i]
	}

	return profit
}
