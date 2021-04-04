//方法一
func isPowerOfTwo(n int) bool { 
	if n==1 {
        return true
    }
    for n!=1 && n!=0 {
        if n&1==1 {
            return false
        }
        n >>= 1
    }
    return false
}


//方法二
func isPowerOfTwo(n int) bool { // 神仙方法
    if n!=0 && n&(n-1) == 0 {
        return true
    }
    return false
}
