package algorithm

func maxArea(height []int) int {
	var max int
	var l, r = 0, len(height) - 1
	for l < r {
		var minHeight int
		if height[l] < height[r] {
			minHeight = height[l]
			l++
		} else {
			minHeight = height[r]
			r--
		}
		var area = minHeight * (r - l + 1)
		max = maxValue(max, area)
	}
	return max
}

func maxValue(i, j int) int {
	if i > j {
		return i
	}
	return j
}
