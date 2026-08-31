// the first solution that comes to my mind is getting the squares of the elements and then returning them after sorting

// broken down solution
// 1.// create a slice with the exact length of nums
	// this is where we will store our final solution

	// 2.// we can then square each element of the array through a loop

	// 3.// then store the squared elemnts in our previous initializzed array

	// 4.// finally i would sort the array and return the solution

func sortedSquares(nums []int) []int {	
	arr:=make([]int, len(nums))

	for i:=0;i<len(nums);i++ {
		sqr:=nums[i]*nums[i]
		arr[i]=sqr

	}	
	sort.Ints (arr[:])
	return arr
}
