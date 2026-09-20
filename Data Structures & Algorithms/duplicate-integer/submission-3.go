// brute force solution would be comparing the elements of the array to each other
// this can be achieved by having 2 loop
// first outer loop and an inner loop fo comparison

func hasDuplicate(nums []int) bool {
    for i:=0;i<len(nums);i++{
        for j:=i+1 ;j<len(nums); j++ {
            if nums[i] ==nums[j] {
                return true
            }
        }
    }
   return false 
}
