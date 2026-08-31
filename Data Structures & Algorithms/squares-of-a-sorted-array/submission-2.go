// for my optimized solution im thinking of implementing a two pointer


// where ill have one pointer to the right and another to the left
//ill then compare the two to find the larger absolute value
// the next step is to square it and insert into an output array 
// we do this in loop in descending order until we have covered all the values
// return the result


func sortedSquares(nums []int) []int {
	res:=make([]int , len(nums))
	l,r:=0, len(nums)-1

	for i:=len(nums)-1;i>=0;i--{
		if nums[l]*nums[l]>nums[r]*nums[r]{
			res[i]=nums[l]*nums[l]
			l++
		}else{
			res[i]=nums[r]*nums[r]
			r--
		}
	}
	return res
}
