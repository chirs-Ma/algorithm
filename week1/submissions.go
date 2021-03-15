func nthUglyNumber(n int) int {
	a, b, c := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		na, nb, nc := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = min(na, nb, nc)
		if dp[i] == na {
			a++
		}
		if dp[i] == nb {
			b++
		}
		if dp[i] == nc {
			c++
		}
	}
	return dp[n-1]
}
func min(a, b, c int) int {
	tmp := a
	if a > b {
		tmp = b
	}
	if tmp > c {
		tmp = c
	}
	return tmp
}