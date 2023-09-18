package main

//func maxProfit(prices []int) int {
//	dayOfBuying := 0
//	totalDays := len(prices)
//	dayOfSelling := totalDays - 1
//	for i := 1; i < totalDays; i++ {
//		if prices[dayOfBuying] > prices[i] {
//			dayOfBuying = i
//		}
//		if prices[dayOfSelling] < prices[totalDays-i] {
//			dayOfSelling = totalDays -i
//		}
//		println(dayOfBuying, dayOfSelling, i, totalDays-i)
//		if (totalDays-i) - i == (totalDays %2) {
//			if dayOfBuying < dayOfSelling && prices[dayOfSelling] > prices[dayOfBuying] {
//				return prices[dayOfSelling] - prices[dayOfBuying]
//			}
//			return 0
//		}
//	}
//	if dayOfBuying < dayOfSelling {
//		return prices[dayOfSelling] - prices[dayOfBuying]
//	}
//	return 0
//}

func maxProfit(prices []int) int {
	buyPrice := prices[0]
	maximumProfit := 0
	for day := 1; day < len(prices); day++ {
		if buyPrice > prices[day] {
			buyPrice = prices[day]
		} else {
			profit := prices[day] - buyPrice
			if profit > maximumProfit {
				maximumProfit = profit
			}
		}
	}
	return maximumProfit
}

func main() {
	//prices := []int{7, 1, 5, 3, 6, 4}
	//prices := []int{1, 4, 2}
	//prices := []int{7, 6, 4, 3, 1}
	//prices := []int{7}
	//prices := []int{2, 1}
	//prices := []int{2, 4, 1}
	//prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	prices := []int{3, 2, 6, 5, 0, 3}

	maxProfit := maxProfit(prices)
	println(maxProfit)
}
