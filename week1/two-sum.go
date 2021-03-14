//首先采取最粗暴的解法双循环法，先遍历一遍数组然后在遍历中在遍历一遍数组，然后将两个遍历出来的值相加判断是否符合要求。
//该解法的进行了两次的for循环所以时间复杂度为O(n^2)
func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j, v2 := range nums {
			if i != j && v+v2 == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

//第二种是采用map进行求解，同样遍历一遍数组，在遍历数组的时候构建map[v]=j的元素，然后判断map[range-v]是否存在。
//该解法只进行了一次for循环因此时间复杂度为O(n)
func twoSum(nums []int, target int) []int {
	mymap := map[int]int{}
	for i, v := range nums {
		if j, ok := mymap[target-v]; ok {
			return []int{i, j}
		}
		mymap[v] = i
	}
	return []int{}
}