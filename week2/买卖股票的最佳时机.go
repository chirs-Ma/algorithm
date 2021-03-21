//采用贪心算法，只要后面的价格大于前面的就可以买入
func maxProfit(prices []int) int {
    profit := 0 
    for i := 0; i < len(prices)-1; i++ {
        if prices[i+1] > prices[i] {
            profit += prices[i+1]-prices[i]
        }
    }
    return profit
}