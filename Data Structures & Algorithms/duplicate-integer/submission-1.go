func hasDuplicate(nums []int) bool {
    mp := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, ok := mp[num]; ok {
			return ok
		}
			mp[num] = struct{}{}
	}
	return false
}
