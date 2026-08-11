func maxArea(heights []int) int {

	maxx := 0
	l, r := 0, len(heights)-1

	for l < r {
		width := r - l
		height := min(heights[l], heights[r])
		area := width * height
		if area > maxx {
			maxx = area
		}

		if heights[l] >= heights[r] {
			r--
		} else {
			l++
		}
	}
	return maxx
}
