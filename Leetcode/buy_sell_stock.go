package Leetcode

// You are given an array prices where prices[i] is the price of a given stock on the ith day.
//
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
//
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
func buyAndSellStock1(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	minPrice := prices[0]
	profit := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		if profit < price-minPrice {
			profit = price - minPrice
		}

	}
	return profit
}

func buyAndSellStock2(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
