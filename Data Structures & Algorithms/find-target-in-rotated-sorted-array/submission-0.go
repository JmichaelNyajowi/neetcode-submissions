// loop through the array
// compare the values of the array to the target
// if they are equal wereturn the index of the value
// if the value does not exist we return -1

func search(nums []int, target int) int {
	for i:=0;i<len(nums);i++{
		if nums[i]==target{
			return i
		}
	}
	return -1

}
