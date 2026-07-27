
func groupAnagrams(strs []string) [][]string {
	resMP := make(map[[26]int][]string, len(strs))
	for _, str := range strs {
		var arr [26]int
		for _, chr := range str {
			arr[chr-'a']++
		}
		resMP[arr] = append(resMP[arr], str)
	}

	res := make([][]string, 0, len(resMP))
	for _, strings := range resMP {
		res = append(res, strings)
	}

	return res
}
