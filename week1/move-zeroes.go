//采用双针针法,j提供可替换的坑位，循环数组进行数据的替换
func moveZeroes(nums []int) {
	i, j := 0, 0
	length := len(nums)
	for i < length {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
		i++
	}
}

//采用非零项覆盖数组头的方式，同样j提供不为0的坑位，剩下的就都为0
func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
}
