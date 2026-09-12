/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	hash:=make(map[*ListNode]bool)
	cur:=head

	for cur!=nil{
		if hash[cur]{
			return true
		}
		hash[cur]=true
		cur=cur.Next
	}
	return false
    
}
