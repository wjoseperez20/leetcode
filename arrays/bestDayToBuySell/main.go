package main

import "fmt"

func maxProfitMedium(prices []int) int {
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		// If the price on the current day is higher than the previous day,
		// we can make a profit by buying and selling on the same day
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
}
func maxProfit_(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	maxProfit := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {

		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			profit := prices[i] - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}

	}

	return maxProfit
}

func maxProfit(prices []int) int {
	var profit int

	if len(prices) <= 1 {
		return profit
	}

	var bestDayBuy, bestDaySell, check int
	bestDayBuy = prices[0]
	bestDaySell = prices[1]
	for idx := 2; idx < len(prices); idx++ {
		// Buy
		if prices[idx-1] <= bestDayBuy {
			bestDayBuy = prices[idx-1]
			check = idx - 1
		}

		// Sell
		if prices[idx] > bestDaySell && idx > check || idx-1 == check {
			bestDaySell = prices[idx]
		}
	}

	profit = bestDaySell - bestDayBuy
	if profit <= 0 {
		return 0
	}

	return profit
}

func main() {
	testCases := []struct {
		prices []int
		profit int
	}{
		{[]int{2, 1, 2, 1, 0, 1, 2}, 2},
		{[]int{2, 1, 2, 1, 0, 1, 2, 3}, 3},
		{[]int{2, 1, 2, 1, 0, 1, 2, 3, 4}, 4},
		{[]int{2, 1, 2, 1, 0, 1, 2, 3, 4, 5}, 5},
	}

	for _, tc := range testCases {
		if profit := maxProfit(tc.prices); profit != tc.profit {
			fmt.Printf("maxProfit(%v) should have been %d but was %d\n", tc.prices, tc.profit, profit)
		}
	}
}
