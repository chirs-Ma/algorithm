//采用回溯的方式进行解答
func subsets(nums []int) [][]int {
    result := make([][]int, 0)
    subsetsBT(&result, nums, []int{}, 0)
    return result
}
func subsetsBT(result *[][]int, nums []int, temp []int, start int) {
    //此处深拷贝temp，避免回溯的时候temp被修改后会影响之前保存的结果
    c := make([]int, len(temp))
    copy(c, temp)
    *result = append(*result, c)

    for i := start; i < len(nums); i++ {
        temp = append(temp, nums[i])
        subsetsBT(result, nums, temp, i+1)//不包含重复值
        temp = temp[:len(temp)-1]
    }
}