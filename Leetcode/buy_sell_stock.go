package Leetcode

// buyAndSellStock1 You are given an array prices where prices[i] is the price of a given stock on the ith day.
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

// buyAndSellStock2 You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
//
// On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.
//
// Find and return the maximum profit you can achieve.
//
// Example 1:
//
// Input: prices = [7,1,5,3,6,4]
// Output: 7
// Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
// Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
// Total profit is 4 + 3 = 7.
// Example 2:
//
// Input: prices = [1,2,3,4,5]
// Output: 4
// Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
// Total profit is 4.
// Example 3:
//
// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.
//
// Constraints:
//
// 1 <= prices.length <= 3 * 104
// 0 <= prices[i] <= 104
func buyAndSellStock2(prices []int) int {

	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

//func buyAndSellStock3(prices []int) int {
//
//	profit := 0
//	for i := 1; i < len(prices); i++ {
//		if prices[i]-prices[i-1] > 0 {
//			profit += prices[i] - prices[i-1]
//		}
//	}
//	return profit
//}
