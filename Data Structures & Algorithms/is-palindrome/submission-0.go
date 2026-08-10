
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {

		for l < r && !isAlphaNum(s[l]) {
			l++
		}
		for r > l && !isAlphaNum(s[r]) {
			r--
		}

		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		l++
		r--

	}
	return true
}

func isAlphaNum(s byte) bool {
	return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || (s >= '0' && s <= '9')
}

