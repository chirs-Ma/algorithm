func relativeSortArray(arr1 []int, arr2 []int) []int {
    lastIndex := 0
    for i := 0; i < len(arr2); i++ {
        for j := lastIndex; j < len(arr1); j++ {
            if arr1[j] == arr2[i] {
                arr1[lastIndex],arr1[j] = arr1[j],arr1[lastIndex]
                lastIndex++
            }
        }
    }
    sort.Ints(arr1[lastIndex:])
    return arr1
}