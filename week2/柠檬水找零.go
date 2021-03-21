//10只能使用5找零，20可以使用10+5或者单独5找零
func lemonadeChange(bills []int) bool {
    fiveCount := 0
    tenCount := 0
    for _,value := range(bills) {
        if value == 5 {
            fiveCount++
        }
        if value == 10 {
            if fiveCount < 1 {
                 return false
            } else {
                fiveCount--
                tenCount++
            } 
        } 
        if value == 20 {
            if tenCount > 0 {
                if fiveCount < 1 {
                    return false
                } else {
                    fiveCount--
                    tenCount--
                }
            } else {
                if fiveCount < 3 {
                    return false
                } else {
                    fiveCount -= 3
                }
            }
        }
    }
    return true;
}