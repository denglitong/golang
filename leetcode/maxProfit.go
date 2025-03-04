package leetcode

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock
func maxProfit(prices []int) int {
	// f(i) - max profit
	// f(i+1)
	// res = max()
	res := 0
	for i := 0; i < len(prices); i++ {
		profit := max(
			profitOnBuy(prices, i),
			0,
		)
		res = max(res, profit)
	}
	return res
}

func profitOnBuy(prices []int, i int) int {
	res := 0
	if i == len(prices)-1 {
		return res
	}
	buyPrice := prices[i]
	for j := i + 1; j < len(prices); j++ {
		if prices[j] > buyPrice {
			res = max(res, prices[j]-buyPrice)
		}
	}
	//fmt.Println(i, res, prices)
	return res
}

func maxProfitV2(prices []int) int {
	// f(i) - max profit
	// f(i+1)
	// res = max()
	res, n := 0, len(prices)
	rightMaxs := make([]int, n)
	rightMax := 0
	for i := n - 1; i >= 0; i-- {
		if prices[i] > rightMax {
			rightMax = prices[i]
		}
		rightMaxs[i] = rightMax
	}
	for i := 0; i < len(prices); i++ {
		res = max(res, profitOnBuyV2(prices, i, rightMaxs))
	}
	return res
}

func profitOnBuyV2(prices []int, i int, rightMaxs []int) int {
	if i == len(prices)-1 {
		return 0
	}
	buyPrice := prices[i]
	maxSellPrice := rightMaxs[i+1]
	if maxSellPrice > buyPrice {
		return maxSellPrice - buyPrice
	}
	//fmt.Println(i, res, prices)
	return 0
}
