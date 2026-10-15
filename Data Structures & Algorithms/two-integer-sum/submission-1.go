func twoSum(nums []int, target int) []int {
    if len(nums) < 2 {
		return nil
	}
	result := make([]int, 0, 2)
	seen := make(map[int]int)
	for idx, val := range nums {
		if seen_value, ok := seen[val]; ok {
			result = append(result, seen_value, idx)
			break
		}
		seen[target - val] = idx
	}
	return result
}
