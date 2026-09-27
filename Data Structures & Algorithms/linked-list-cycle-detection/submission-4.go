/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil { 
		slow = slow.Next
		fast = fast.Next
		if fast != nil { 
			fast = fast.Next
			if slow == fast { 
				return true
			}
		}
	}

	return false
    
}
