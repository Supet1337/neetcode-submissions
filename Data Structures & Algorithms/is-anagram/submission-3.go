
func isAnagram(s string, t string) bool {
	mpS := make(map[string]int, len(s))
	mpT := make(map[string]int, len(t))

	for _, char := range s {
		mpS[string(char)]++
	}

	for _, char := range t {
		mpT[string(char)]++
	}

	if len(mpS) != len(mpT) {
		return false
	}

	for chr, cnt := range mpS {
		if _, ok := mpT[chr]; !ok {
			return false
		}

		if cnt != mpT[chr] {
			return false
		}
	}

	return true
}
