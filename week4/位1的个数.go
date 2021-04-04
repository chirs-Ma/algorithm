//方法一循环检查
func hammingWeight(num uint32) (ones int) {
    for i := 0; i < 32; i++ {
        if 1<<i&num > 0 {
            ones++
        }
    }
    return
}

//位运算优化
func hammingWeight(num uint32) (ones int) {
    for ; num > 0; num &= num - 1 {
        ones++
    }
    return
}

