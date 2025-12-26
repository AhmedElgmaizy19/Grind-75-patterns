package TwoPointer

import "sort"


func ThreeSumValue(nums []int,target int) bool {
	sort.Ints(nums)
	for i:=0; i<len(nums)-2;i++ {
		low := i+1
		high := len(nums)-1

		for low <  high {
			trible := nums[i] + nums[low] + nums[high]
			if target == trible{
			return true
			}else if target > trible {
				low++
			}else {
				high--
			}
		}
	}
	return false
}