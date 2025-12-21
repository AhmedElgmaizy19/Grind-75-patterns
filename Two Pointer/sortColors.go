package TwoPointer



func sortColor(colors[]int) []int {
	start , end , current := 0 , len(colors) - 1 , 0

	for current <= end {
		switch colors[current] {
		case 0:
			if colors[start]!= 0 {
				colors[start], colors[current] = colors[current], colors[start]
			}
			start++
			current++
		case 1:
			current++
		case 2:
			if colors[end] != 2 {
				colors[end] , colors[current] = colors[current] , colors[end]
			}
			end--
		}
		
	}
	return colors



}