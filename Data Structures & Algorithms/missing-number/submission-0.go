func missingNumber(nums []int) int {
	mp := make(map[int]struct{})

	for _, v := range nums {
		mp[v] = struct{}{}
	}

	for i := range nums {
		if _, ok := mp[i]; !ok {
			return i
		}
	}
	return len(nums)
}
