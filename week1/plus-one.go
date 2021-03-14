//一个数加一分为这个数为9和不为9两种情况，当这个数为9时需要进位;当出现9,99...特殊情况时需要再最前面加一个1
func plusOne(digits []int) []int {
	res := []int{1}
	lenght := len(digits)
	for i := lenght - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append(res, digits...)
}