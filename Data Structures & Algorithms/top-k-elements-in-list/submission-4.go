func topKFrequent(nums []int, k int) []int {
	maxFreq := -1
	mpNumFreq := make(map[int]int)
	for _, num := range nums {
		mpNumFreq[num]++
		if mpNumFreq[num] > maxFreq {
			maxFreq = mpNumFreq[num]
		}
	}

	freqArr := make([][]int, maxFreq+1)
	for num, freq := range mpNumFreq {
		freqArr[freq] = append(freqArr[freq], num)
	}

	res := make([]int, 0, k)
	fmt.Println(freqArr)
	for i := len(freqArr) - 1; i >= 0; i-- {
		need := k - len(res)
		if need == 0 {
			return res
		}

		if len(freqArr[i]) <= need {
			res = append(res, freqArr[i]...)
		} else {
			res = append(res, freqArr[i][:need]...)
		}

	}
	return res
}

