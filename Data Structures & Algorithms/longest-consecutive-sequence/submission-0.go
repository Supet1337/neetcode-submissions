func longestConsecutive(nums []int) int {
	mp := make(map[int]struct{})
	for _, num := range nums {
		mp[num] = struct{}{}
	}

	var maxx int
	for num := range mp {
		if _, ok := mp[num-1]; ok {
			continue
		}
		nextNum := num + 1
		localMax := 1
		for {
			if _, ok := mp[nextNum]; ok {
				localMax++
				nextNum += 1
			} else {
				break
			}
		}
		if localMax > maxx {
			maxx = localMax
		}
	}

	return maxx
}