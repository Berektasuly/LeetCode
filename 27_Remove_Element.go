func removeElement(nums []int, val int) int {
    var newarr []int
    value := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            newarr = append(newarr, nums[i])
            value += 1
        }
    }
    copy(nums, newarr)
    return value
}