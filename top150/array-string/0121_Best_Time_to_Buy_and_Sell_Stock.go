package arraystring

func maxProfit(prices []int) int {
	// need to find the largest jump or delta
	// set min to first element
	// if you find min lesser than current then set min to that
	// at each iteration calculate profit and store max profit

	n := len(prices)
	if n == 1 {
		return 0
	}

	profit := 0
	mini := prices[0]

	for i := 1; i < n; i++ {
		if prices[i] < mini {
			mini = prices[i]
			continue
		}
		if prices[i]-mini > profit {
			profit = prices[i] - mini
		}
	}

	return profit
}
