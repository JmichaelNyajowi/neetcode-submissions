//BRUTE FORCE SOLUTION
// scan through the entire array
// track the smallest value 
// return the smallest value
// time complexity-o(n)
// space complexity-o(1)

func findMin(nums []int) int {
	minVal:=nums[0]
	for _, val:= range nums{
		if val<minVal{
			minVal=val
		}
	}
	return minVal

}
