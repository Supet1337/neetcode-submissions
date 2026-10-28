func maxProfit(prices []int) int {
	minn, maxx, maxProf := 101, -1, 0
	for _, price := range prices {
		if price < minn {
			if maxx-minn > maxProf {
				maxProf = maxx - minn
			}

			minn = price
			maxx = -1
			continue
		}
		if price > maxx {
			maxx = price
			if maxx-minn > maxProf {
				maxProf = maxx - minn
			}
		}
	}

	return maxProf
}