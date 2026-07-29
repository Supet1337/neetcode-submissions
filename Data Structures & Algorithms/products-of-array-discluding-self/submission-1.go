func productExceptSelf(nums []int) []int {
	prefixes := make([]int, len(nums))
	coef := 1
	for i, num := range nums {
		coef = coef * num
		prefixes[i] = coef
	}
	//fmt.Println(prefixes)

	suffixes := make([]int, len(nums))
	coef = 1
	for i := len(nums) - 1; i >= 0; i-- {
		coef = coef * nums[i]
		suffixes[i] = coef
	}
	//fmt.Println(suffixes)

	res := make([]int, len(nums))
	for i := range nums {
		prefi, sufi := i-1, i+1
		if prefi < 0 {
			prefi = i
		}
		if sufi >= len(nums) {
			sufi = len(nums) - 1
		}

		if i == len(nums)-1 {
			res[i] = prefixes[prefi]
		} else if i == 0 {
			res[i] = suffixes[sufi]
		} else {
			res[i] = prefixes[prefi] * suffixes[sufi]
		}
		//fmt.Println("res[", i, "]", "=", "prefixes[", prefi, "]", "*", "suffixes[", sufi, "]")
	}

	return res
}