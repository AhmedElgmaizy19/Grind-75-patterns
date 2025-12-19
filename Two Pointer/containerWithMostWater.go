package TwoPointer

import "math"

func maxArea(height []int) int {
	start := 0
	end := len(height) - 1
	maxArea := 0

	for start < end {
		width := end - start
		area := int(math.Min(float64(height[start]), float64(height[end]))) * width
		maxArea = int(math.Max(float64(maxArea), float64(area)))
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}
	
	return maxArea
}
