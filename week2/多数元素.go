//解法一：使用哈希实现
func majorityElement(nums []int) int {
    n := len(nums)
    myMap := make(map[int]int)
    for i := 0;i <= n;i++ {
        myMap[nums[i]]++
        if myMap[nums[i]] > n/2 {
            return nums[i]
        }
    }
    return -1
}
//解法二：参考题解中的摩尔投票实现
func majorityElement(nums []int) int {
    count,major := 0,0
    for _,num := range(nums) {
        if count == 0 {
            major = num
        }
        if major == num {
            count++
        } else {
            count--
        }
    }
    return major
}