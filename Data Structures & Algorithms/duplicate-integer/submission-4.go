// we can optimize this to o(nlog n)time complexity using sorting 
// first we sort the array
// then run a loop while comparing every value with the next value close to it
// and if we have a duplicate we return true

func hasDuplicate(nums []int) bool {
    sort.Ints(nums)
    for i:=1;i<len(nums);i++{
        if nums[i]==nums[i-1]{
            return true
        }

    }
    return false
    
}
