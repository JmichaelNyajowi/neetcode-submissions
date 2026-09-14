/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//  for optimization we can use a 2 pointer(slow and fast)
// IMPLEMENTATION
// iniatialize the 2 pointers
// slow
// fast
// if fast pointer is not equal to nil and fast.Next is not equal to nil, slow=slow.Next and fast=fast.Next.Next
// if slow==fast return true

func hasCycle(head *ListNode) bool {
	slow:=head
	fast:=head

	for fast != nil && fast.Next != nil {
		slow =slow.Next
		fast =fast.Next.Next
		if slow==fast{
			return true
		}
		
	}
    return false
}
