func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	max := 0

	for l < r {
		// 1. Calculate the current area
		area := (r - l) * least(height[r], height[l])
		if area > max {
			max = area
		}

		// 2. Simply move the pointer pointing to the shorter line
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}
func least(n1, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}
/// 1 7 least elemt is bigger thanseicns least elemt 
// 6 * 6 =36  6*6 