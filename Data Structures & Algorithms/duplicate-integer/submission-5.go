// we can also optimize the solution using a hashmap to o(n) time and space complexity if we use extra memory
// we iterate through the array using a loop
// for each iteration we confirm if the value we are at is in our hashmap already, if yes we retun true, if not we add it into the hash and proceed with the iteration


func hasDuplicate(nums []int) bool {
    seen := make(map[int]bool)
    for _, num := range nums {
        if seen[num] {
            return true
        }
        seen[num] = true
    }
    return false
}