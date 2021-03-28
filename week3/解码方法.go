func numDecodings(s string) int {
    if s[0] == '0' {
        return 0
    }
    if len(s) == 1 {
        return 1
    }
    dp := make([][]int, len(s))
    for i := range dp {
        dp[i] = make([]int, 2)
    }
    dp[0][0] = 1
    dp[0][1] = 0
    for i := 1; i < len(s); i++ {
        num := s[i-1:i+1]
        if i == 1 {
            if s[i] == '0' {
                dp[i][0] = 0
            } else {
                dp[i][0] = 1
            }
            if num >= "1" && num <= "26" {
                dp[i][1] = 1
            } else {
                dp[i][1] = 0
            }
            continue
        }
        if s[i] == '0' {
            dp[i][0] = 0
        } else {
            dp[i][0] = dp[i-1][0] + dp[i-1][1]
        }
        if num >= "1" && num <= "26" {
            dp[i][1] = dp[i-2][0] + dp[i-2][1]
        } else {
            dp[i][1] = 0
        }
    }
    return dp[len(s)-1][0] + dp[len(s)-1][1]
}

