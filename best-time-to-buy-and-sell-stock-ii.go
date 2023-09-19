package main

func maxProfit(prices []int) int {
	buyPrice := prices[0]
	maximumProfit := 0
	for day := 1; day < len(prices); day++ {
		if buyPrice > prices[day] {
			buyPrice = prices[day]
		}
		profit := prices[day] - buyPrice
		if profit > 0 {
			maximumProfit += profit
			buyPrice = prices[day]
		}
	}
	return maximumProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	//prices := []int{1, 2, 3, 4, 5}
	//prices := []int{7, 6, 4, 3, 1}
	//prices := []int{2, 1, 2, 0, 1}

	maxProfit := maxProfit(prices)
	println(maxProfit)
}
