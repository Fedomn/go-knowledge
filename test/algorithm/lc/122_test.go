package lc

func maxProfit122(prices []int) int {
	res := 0
	for i := 0; i <= len(prices)-2; i++ {
		if prices[i] < prices[i+1] {
			res += prices[i+1] - prices[i]
		}
	}
	return res
}
