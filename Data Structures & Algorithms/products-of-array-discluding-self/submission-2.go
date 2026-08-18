func exceptSelf(nums []int) []int{
    output:=make([]int,len(nums))

    for i:=0;i<len(nums);i++{
        prod:=1
        for j:=0; j<len(nums); j++{
            if j==i{
                continue
            }
            
            prod*=nums[j]

        }
        output[i]=prod
    }
    return output

}






























func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))
    for i := range res {
        res[i] = 1
    }

    prefix := 1
    for i := 0; i < len(nums); i++ {
        res[i] = prefix
        prefix *= nums[i]
    }

    postfix := 1
    for i := len(nums) - 1; i >= 0; i-- {
        res[i] *= postfix
        postfix *= nums[i]
    }

    return res
}
