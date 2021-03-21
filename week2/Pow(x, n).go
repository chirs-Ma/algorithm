//采用暴力求解,代码略过


//采用分治的思想
func myPow(x float64, n int) float64 {
    if n > 0 {
        return fastPow(x, n)
    } else {
        return 1.0/fastPow(x, n)
    }
}
func fastPow (x float64, n int) float64 {
    if n == 0 {
        return 1.0
    }
    subResult := fastPow(x,n/2)
    if n %2 == 0 {
        return subResult * subResult
    } else {
        return subResult * subResult * x
    }

}

