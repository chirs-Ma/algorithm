//栈
func longestValidParentheses(s string) int {
    maxAns := 0
    stack := []int{}
    stack = append(stack, -1)
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            stack = append(stack, i)
        } else {
            stack = stack[:len(stack)-1]
            if len(stack) == 0 {
                stack = append(stack, i)
            } else {
                maxAns = max(maxAns, i - stack[len(stack)-1])
            }
        }
    }
    return maxAns
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

//动态规划
func longestValidParentheses(s string) int {
    max := 0
    dp := make([]int, len(s))
    for i,l := 1, len(s);i<l;i++{
        if s[i] == ')'{
            if k := i-dp[i]-1; k >= 0 && s[k] == '('{
                m := dp[i]+ 2 + dp[k]
                if i + 1 < l {
                    dp[i+1] = m
                }
                if m > max{
                    max = m
                }
            }
        }
    }
    return max
}

