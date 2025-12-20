package TwoPointer

//  naive solution O(n^2) (Not optimal)
// func productExceptSelf(nums[]int) []int {
// 	n := len(nums)
// 	result := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		pro := 1
// 		for j := 0; j < n; j++ {
// 			if i != j {
// 				pro *= nums[j]
// 			}
// 		}
// 		result[i] = pro
// 	}
// 	return result
// }




// optimal solution O(n)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int,n)
	for i:= range(res){
		res[i] = 1
	} 

	start , end := 0 , n-1
	left_proudct ,right_proudct := 1,1
	for start < n && end >= 0 {
		res[start]*=left_proudct
		res[end]*=right_proudct
		left_proudct *= nums[start]
		right_proudct *= nums[end]
		start++
		end--
	} 
	return res
}