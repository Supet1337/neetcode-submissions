func lengthOfLongestSubstring(s string) int {
	mp := make(map[rune]struct{})
	l, r, diff, maxxD := 0, 0, 0, 0

	for r < len(s) {
		if _, ok := mp[rune(s[r])]; !ok {
			mp[rune(s[r])] = struct{}{}
			diff = r - l + 1
			if diff > maxxD {
				maxxD = diff
			}
			r++
			continue
		} else {
			delete(mp, rune(s[l]))
			l++
		}

	}

	return maxxD
}