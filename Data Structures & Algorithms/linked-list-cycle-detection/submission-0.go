/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	for head != nil { 
		
		 head = head.Next
		 fast = head.Next

		 if head.Val == fast.Val { 
			return true
		 }
		 

	}

	return false 
}
