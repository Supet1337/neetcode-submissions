func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, num := range nums {
		mp[target-num] = i
	}

	for i, num := range nums {
		if dem, ok := mp[num]; ok && dem != i {
			return []int{i, dem}
		}
	}
	return []int{}
}
