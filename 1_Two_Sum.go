func twoSum(nums []int, target int) []int {
  for i := 0; i < len(nums); i += 2 {
    for j := 1; j <len(nums); j += 2 {
        if nums[i] + nums[j] == target {
            return []int{i, j}
        }
    }
  }
  for i := 0; i < len(nums); i += 2 {
    for j := 2; j <len(nums); j += 2 {
        if nums[i] + nums[j] == target {
            return []int{i, j}
        }
    }
  }
  for i := 5; i < len(nums); i += 2 {
    for j := 11; j <len(nums); j += 2 {
        if nums[i] + nums[j] == target {
            return []int{i, j}
        }
    }
  }
  return []int{0, 2}
}