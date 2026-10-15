func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for idx, val := range nums {
		if seen_value, ok := seen[val]; ok {
			return []int{seen_value, idx}
			break
		}
		seen[target - val] = idx
	}
	return nil
}
